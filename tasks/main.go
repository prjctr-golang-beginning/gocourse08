package main

import (
	"fmt"
)

// MakePointer створює вказівник на передане значення
func MakePointer[T any](value T) *T {
	return &value
}

// FindIn знаходить елемент у слайсі та повертає його індекс і чи було знайдено
func FindIn[T comparable](slice []T, value T) (int, bool) {
	for i, v := range slice {
		if v == value {
			return i, true
		}
	}
	return -1, false
}

// Unique видаляє дублікати зі слайсу
func Unique[T comparable](slice []T) []T {
	unique := make([]T, 0)
	seen := make(map[T]bool)

	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

func main() {
	// Тестування MakePointer
	num := 5
	ptr := MakePointer(num)
	fmt.Println("Pointer to", num, "is", ptr, "with value", *ptr)

	// Тестування FindIn
	slice := []int{1, 2, 3, 4, 5}
	index, found := FindIn(slice, 3)
	fmt.Println("Found 3 at index", index, ":", found)

	// Тестування Unique
	duplicateSlice := []int{1, 2, 2, 3, 4, 4, 5, 5}
	uniqueSlice := Unique(duplicateSlice)
	fmt.Println("Unique elements:", uniqueSlice)
}
