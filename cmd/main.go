package main

import (
	"context"
	"racecondition/auditor"
	"time"
)

type myRepository struct {
}

type Value string

func (v Value) Value() any {
	return string(v)
}

func (r *myRepository) CreateMany(_ context.Context, _ []auditor.Valuable) (int, error) {
	return 4, nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	a := auditor.New(new(myRepository), ctx)
	val1 := Value(`Some value 1`)
	val2 := Value(`Some value 2`)

	a.Update(val1)
	a.Update(val2)

	time.Sleep(time.Second * 2)

	a.Update(val1)
	a.Update(val2)

	time.Sleep(time.Second * 2)

	cancel()

	time.Sleep(time.Second * 2)

	// exercises
}
