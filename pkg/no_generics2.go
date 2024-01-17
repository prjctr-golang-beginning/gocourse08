package pkg

type Student struct {
	Id   int
	Name string
}

func (s *Student) SetId(id int) {
	s.Id = id
}

type Employee struct {
	Id     int
	Name   string
	Salary float64
}

func (s *Employee) SetId(id int) {
	s.Id = id
}

type Product struct {
	Id    int
	Name  string
	Price float64
}

func (s *Product) SetId(id int) {
	s.Id = id
}

type Event struct {
	Id   int
	Name string
}

func (s *Event) SetId(id int) {
	s.Id = id
}

type Vehicle struct {
	Id    int
	Model string
}

func (s *Vehicle) SetId(id int) {
	s.Id = id
}

func SetStudentId(s Student, id int) Student {
	s.SetId(id)
	return s
}

func SetEmployeeId(e Employee, id int) Employee {
	e.SetId(id)
	return e
}

func SetProductId(p Product, id int) Product {
	p.SetId(id)
	return p
}

func SetEventId(ev Event, id int) Event {
	ev.SetId(id)
	return ev
}

func SetVehicleId(v Vehicle, id int) Vehicle {
	v.SetId(id)
	return v
}
