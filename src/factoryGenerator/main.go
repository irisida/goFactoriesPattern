package main

import "fmt"

type Employee struct {
	Name, Position string
	Income         int
}

// functional factory pattern
// seemingly more idiomatic to Go
func NewEmployeeFactory(position string, income int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, income}
	}
}

// Alternate approach
type EmployeeFactory struct {
	Position string
	Income   int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.Income}
}

func NewEmployeeFactoryTwo(position string, income int) *EmployeeFactory {
	return &EmployeeFactory{position, income}
}

func main() {

	// using functional approach
	developerFactory := NewEmployeeFactory("developer", 36000)
	managerFactory := NewEmployeeFactory("Manager", 95000)

	dev := developerFactory("Johnny Five")
	manager := managerFactory("Fat Controller")

	fmt.Println(*dev)
	fmt.Println(*manager)

	// using alternate approach
	bossFactory := NewEmployeeFactoryTwo("CEO", 100000)
	boss := bossFactory.Create("Cheese")
	fmt.Println(*boss)
}
