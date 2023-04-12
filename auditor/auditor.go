package auditor

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Valuable interface {
	Value() any
}

const (
	flushPeriod = 2 * time.Second
	flushMax    = 100

	flushTypePeriod = iota
	flushTypeMax
	flushTypeLast
)

type Repository interface {
	CreateMany(context.Context, []Valuable) (int, error)
}

func New(r Repository, ctx context.Context) *Auditor {
	res := &Auditor{
		repository: r,
		flush:      make(chan int),
		stop:       make(chan struct{}),
		ticker:     time.NewTicker(flushPeriod),
	}

	res.autoFlush(ctx)

	return res
}

type Auditor struct {
	repository Repository
	mu         sync.RWMutex
	_entities  []Valuable
	ticker     *time.Ticker
	flush      chan int
	stop       chan struct{}
	stopped    bool
}

func (a *Auditor) isStopped() bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.stopped
}

func (a *Auditor) autoFlush(ctx context.Context) {
	go func() { // init ticker
		for {
			select {
			case <-a.ticker.C:
				if !a.isStopped() {
					a.flush <- flushTypePeriod
				}
			case <-a.stop:
				a.ticker.Stop()
				fmt.Println(`Ticker stopped`)
				return
			}
		}
	}()

	go func() { // init flusher
		for {
			select {
			case tp := <-a.flush:
				a.Flush(tp)
			case <-a.stop:
				fmt.Println(`Flusher stopped`)
				return
			}
		}
	}()

	go func() { // init context listener
		select {
		case <-ctx.Done():
			fmt.Println(`Context is Done`)
			wg := &sync.WaitGroup{}
			wg.Add(1)
			a.Stop(wg)
			wg.Wait()
			return
		}
	}()
}

func (a *Auditor) Flush(fType int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	entitiesLen := len(a._entities)
	if entitiesLen == 0 {
		return
	}

	affected, err := a.repository.CreateMany(context.Background(), a._entities[0:entitiesLen])
	if err != nil {
		fmt.Printf(`auditor didn't save events. Error: %s. Type: %s`, err, flushType(fType))
	} else {
		log.Printf(`Auditor flush events. Num: %d. Type: %s.`, affected, flushType(fType))
	}

	a._entities = a._entities[entitiesLen:]
}

func (a *Auditor) lastFlush() {
	close(a.flush)
	a.Flush(flushTypeLast)
}

func (a *Auditor) Stop(wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Stop function", r)
		}
	}()

	close(a.stop)

	a.mu.Lock()
	a.stopped = true
	a.mu.Unlock()

	a.lastFlush()
	wg.Done()

	log.Println("---- Auditor log flushed and stopped")
}

func (a *Auditor) Update(subj any) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Update function", r)
		}
	}()

	ent, ok := subj.(Valuable)
	if !ok {
		fmt.Printf(`subject for Auditor is not Valuable type: Actual type: %s`, fmt.Sprintf("%T", subj))
	}

	if a.isStopped() {
		//log.Errorw(`Auditor is trying updates after it had been stopped`, `entity`, ent)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a._entities = append(a._entities, ent)
	a.triggerFlushMax()
}

func (a *Auditor) triggerFlushMax() {
	if flushMax == len(a._entities) {
		a.ticker.Reset(flushPeriod)
		a.flush <- flushTypeMax
	}
}

func flushType(tp int) string {
	switch tp {
	case flushTypePeriod:
		return `period`
	case flushTypeMax:
		return `max`
	case flushTypeLast:
		return `last`
	default:
		return `undefined`
	}
}
