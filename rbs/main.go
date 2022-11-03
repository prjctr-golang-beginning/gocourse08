// Rough brush stroke
package main

import "fmt"

// 1
type Product1 struct {
	name  string
	price float64
}
type Shipment1 struct {
	products []Product1
	address  string
}

func (s *Shipment1) Ship() {
	fmt.Printf(`Products shipped`)
}

// 2
type Product2 struct {
	name     string
	price    float64
	discount float64
}
type Shipment2 struct {
	products []Product2
	address  Address2
}
type Address2 struct {
	country string
	city    string
	address string
}

func (s *Shipment2) setProducts(p []Product2) {
	s.products = p
}
func (s *Shipment2) setAddress(a Address2) {
	s.address = a
}
func (s *Shipment2) Ship() {
	fmt.Printf(`Products shipped`)
}

// 3
type Product3 struct {
	name  string
	price float64
}
type Shipment3 struct {
	products []Product3
	address  Address3
}
type Address3 struct {
	country string
	city    string
	address string
}
type State struct {
	step int
}

func (s *State) Check() bool {
	s.step++

	if s.step == 3 {
		fmt.Printf(`Products shipped`)
		return true
	} else {
		fmt.Printf(`Products not shipped yet`)
		return false
	}
}

func (s *Shipment3) addProduct(p Product3) {
	s.products = append(s.products, p)
}
func (s *Shipment3) setAddress(a Address3) {
	s.address = a
}
func (s *Shipment3) Ship() *State {
	return &State{}
}

func main() {
	// 1
	shipment1 := &Shipment1{[]Product1{
		{`product 1`, 1},
		{`product 2`, 2},
		{`product 3`, 3},
	}, `Some address`}
	shipment1.Ship()

	// 2
	shipment2 := &Shipment2{}
	shipment2.setProducts([]Product2{
		{`product 1`, 1, 0},
		{`product 2`, 2, 0.1},
		{`product 3`, 3, 0.3},
	})
	shipment2.setAddress(Address2{`Ukraine`, `Kyiv`, `My Home address`})
	shipment2.Ship()

	// 3
	shipment3 := &Shipment3{}
	shipment3.addProduct(Product3{`product 1`, 1})
	shipment3.addProduct(Product3{`product 2`, 2})
	shipment3.addProduct(Product3{`product 3`, 1})
	shipment3.setAddress(Address3{`Ukraine`, `Kyiv`, `My Home address`})
	state := shipment3.Ship()
	for !state.Check() {
		fmt.Printf(`Order completed`)
	}
}
