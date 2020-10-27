package main

import "fmt"

// description interface requires implementation of
// a describe method to adhere
type Description interface {
	Describe()
}

type person struct {
	name string
}

type product struct {
	desc string
}

// implements the Description interface Describe method for person
func (p *person) Describe() {
	fmt.Printf("Name: %s \n", p.name)
}

// implements the Description interface Describe method for product
func (p *product) Describe() {
	fmt.Printf("Product ID: %s \n", p.desc)
}

func NewEntity(isHuman bool, identifier string) Description {
	// validates the type we need to create and
	// returns a pointer to the relevant type
	if isHuman {
		return &person{identifier}
	}
	return &product{identifier}

}

func main() {
	// pass true to define a human
	s := NewEntity(true, "Johnny Five")
	s.Describe()

	// passing false to default to a product
	d := NewEntity(false, "GPU Processor")
	d.Describe()
}
