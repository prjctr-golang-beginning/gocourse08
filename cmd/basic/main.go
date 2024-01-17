package main

import (
	"fmt"
	"github.com/prjctr-golang-beginning/gocourse08/pkg"
)

func main() {
	// Виконання без дженеріків
	fmt.Println(pkg.CalcNum(uint(10), -25.5))

	item1 := pkg.Student{Name: "John"}
	item2 := pkg.Employee{Name: "Alice", Salary: 50000}
	item3 := pkg.Product{Name: "Laptop", Price: 1000}
	item4 := pkg.Event{Name: "Conference"}
	item5 := pkg.Vehicle{Model: "Tesla"}

	fmt.Printf("%T\n", pkg.SetStudentId(item1, 1))
	fmt.Printf("%T\n", pkg.SetEmployeeId(item2, 2))
	fmt.Printf("%T\n", pkg.SetProductId(item3, 3))
	fmt.Printf("%T\n", pkg.SetEventId(item4, 4))
	fmt.Printf("%T\n", pkg.SetVehicleId(item5, 5))

	// Введення в Дженеріки
	fmt.Println(CompareAny(5, 5))          // true
	fmt.Println(CompareAny("Hello", "Hi")) // false

	// Типові Обмеження та Їх Використання
	fmt.Println(SumNumbers([]int{1, 2, 3}))      // 6
	fmt.Println(SumNumbers([]float64{1.5, 2.5})) // 4

	fmt.Println(SumNumbersAny([]uint32{24, 6, 0}))  // 30
	fmt.Println(SumNumbersAny([]float32{4.1, 0.1})) // 4.2

	fmt.Printf("%T\n", GenProcessor[*pkg.Student]{}.Process(&item1, 1))
	fmt.Printf("%T\n", GenProcessor[*pkg.Employee]{}.Process(&item2, 2))
	fmt.Printf("%T\n", GenProcessor[*pkg.Product]{}.Process(&item3, 3))
	fmt.Printf("%T\n", GenProcessor[*pkg.Event]{}.Process(&item4, 4))
	fmt.Printf("%T\n", GenProcessor[*pkg.Vehicle]{}.Process(&item5, 5))

	// Універсальні Структури Даних
	list := GenericList[int]{}
	list.Add(1)
	list.Add(2)
	fmt.Println(list.GetAll()) // [1 2]

	var sum int
	for _, i := range list.GetAll() {
		sum += i
	}
	fmt.Println(sum) // [1 2]
}
