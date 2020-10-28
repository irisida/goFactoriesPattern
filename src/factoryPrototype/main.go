package main

import "fmt"

type Employee struct {
	Name, Position string
	Income         int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 65000}
	case Manager:
		return &Employee{"", "manager", 95000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "Ed"
	fmt.Println(*m)
}
