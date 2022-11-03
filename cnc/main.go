// Coupling and Cohesion
package main

import "fmt"

// coupling
type Product struct {
	name  string
	price float64
}
type Shipment struct {
	products []Product
	address  Address
}
type Address struct {
	country string
	city    string
	address string
}

func (s *Shipment) setProducts(p []Product) {
	s.products = p
}
func (s *Shipment) setAddress(a Address) {
	s.address = a
}
func (s *Shipment) Ship() {
	fmt.Printf(`Products shipped`)
}

// cohesion
type Order struct {
	products []Product
	address  string
	receiver string
	seller   string
	email    string
}

func (o *Order) setProducts(p []Product) {
	o.products = p
}
func (o *Order) setAddress(a string) {
	o.address = a
}
func (o *Order) setReceiver(r string) {
	o.receiver = r
}
func (o *Order) setSeller(s string) {
	o.seller = s
}
func (o *Order) SendEmail() {
	fmt.Printf(`E-mail sent`)
}
func (o *Order) RegisterForShipment(*Shipment) {
	fmt.Printf(`Registered for shipment`)
}

func main() {
	products := []Product{
		{`product 1`, 1},
		{`product 2`, 2},
		{`product 3`, 3},
	}

	// coupling
	shipment := &Shipment{}
	shipment.setAddress(Address{`Ukraine`, `Kyiv`, `My Home address`})
	shipment.setProducts(products)
	shipment.Ship()

	// cohesion
	order := Order{}
	order.setProducts(products)
	order.setAddress(`Some address`)
	order.setReceiver(`Some receiver`)
	order.setSeller(`Some seller`)
	order.SendEmail()
	order.RegisterForShipment(shipment)
}
