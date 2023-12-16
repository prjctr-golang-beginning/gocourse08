package main

import (
	"fmt"
)

// Compare два значення будь-якого типу T
func CompareAny[T comparable /*type constraint*/](a, b T) bool {
	return a == b
}

// SumNumbers сумує масив чисел будь-якого числового типу
func SumNumbers[T int | float64](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// SumNumbers сумує масив чисел будь-якого числового типу
func SumNumbersAny[T Number](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

type Number interface {
	int | int8 | int32 | int64 | uint | uint8 | uint32 | uint64 | float32 | float64
}

// GenericList - це універсальний список
type GenericList[T any] struct {
	elements []T
}

// Add додає елемент у список
func (g *GenericList[T]) Add(element T) {
	g.elements = append(g.elements, element)
}

// GetAll повертає всі елементи списку
func (g *GenericList[T]) GetAll() []T {
	return g.elements
}

func main() {
	// Введення в Дженеріки
	fmt.Println(CompareAny(5, 5))          // true
	fmt.Println(CompareAny("Hello", "Hi")) // false

	// Типові Обмеження та Їх Використання
	fmt.Println(SumNumbers([]int{1, 2, 3}))      // 6
	fmt.Println(SumNumbers([]float64{1.5, 2.5})) // 4

	fmt.Println(SumNumbersAny([]uint32{24, 6, 0}))  // 30
	fmt.Println(SumNumbersAny([]float32{4.1, 0.1})) // 4.2

	// Універсальні Структури Даних
	list := GenericList[int]{}
	list.Add(1)
	list.Add(2)
	fmt.Println(list.GetAll()) // [1 2]
}
