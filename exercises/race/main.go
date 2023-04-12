package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// створюємо контекст для відслідковування завершення гонки
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// створюємо канал для отримання повідомлень про завершення гонки
	finish := make(chan string)

	// створюємо дві горутини, кожна з яких відповідає за одну машину
	go raceCar(ctx, "Машина 1", finish)
	go raceCar(ctx, "Машина 2", finish)

	// чекаємо, поки хоча б одна машина доїде до фінішу
	winner := <-finish

	// виводимо повідомлення про переможця
	fmt.Printf("%s переміг!\n", winner)
}

// функція для змагання машини на гоночному треку
func raceCar(ctx context.Context, name string, finish chan<- string) {
	// встановлюємо рандомну швидкість машини
	speed := time.Duration(rand.Intn(100-50)+50) * time.Millisecond

	// виводимо повідомлення про старт машини
	fmt.Printf("%s стартував зі швидкістю %v\n", name, speed)

	// чекаємо, доки машина не доїде до фінішу або не буде скасовано контекст
	select {
	case <-time.After(speed):
		// машина доїхала до фінішу
		finish <- name
	case <-ctx.Done():
		// гонка скасована
		fmt.Printf("%s не доехав до фінішу\n", name)
	}
}
