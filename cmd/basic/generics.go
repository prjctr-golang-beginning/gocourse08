package main

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

type Identifiable interface {
	SetId(int)
}

type GenProcessor[T Identifiable] struct{}

func (p GenProcessor[T]) Process(item T, id int) T {
	item.SetId(id)
	return item
}
