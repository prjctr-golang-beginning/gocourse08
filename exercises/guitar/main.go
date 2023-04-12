package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	leftHand := make(chan string)
	rightHand := make(chan string)

	// Ліва рука
	go func() {
		for {
			select {
			case note := <-leftHand:
				fmt.Println("Ліва рука грає ноту:", note)
			case <-ctx.Done():
				fmt.Println("Гра припинилась через виключення світла")
				return
			}
		}
	}()

	// Права рука
	go func() {
		for {
			select {
			case note := <-rightHand:
				fmt.Println("Права рука затискає ноту:", note)
			case <-ctx.Done():
				fmt.Println("Гра припинилась через виключення світла")
				return
			}
		}
	}()

	// Симулюємо гру на гітарі
	for i := 0; i < 10; i++ {
		note := generateRandomNote()
		fmt.Println("Граємо ноту:", note)

		// Ліва рука щіпає струни
		leftHand <- note

		// Права рука затискає струни
		rightHand <- note

		time.Sleep(500 * time.Millisecond)
	}

	// Виключення світла
	cancel()
	time.Sleep(500 * time.Millisecond)
}

// Генерує випадкову ноту
func generateRandomNote() string {
	notes := []string{"C", "D", "E", "F", "G", "A", "B"}
	return notes[rand.Intn(len(notes))]
}
