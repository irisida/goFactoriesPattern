package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	CanVote bool
}

// factory function is responsible for creating out
// object in a wholesale manner. It creates and
// validates, returning a completed object.
func NewPerson(name string, age int) *Person {
	if age > 18 {
		return &Person{name, age, true}
	}
	return &Person{name, age, false}

}

func main() {
	p1 := NewPerson("Shug", 40)
	p2 := NewPerson("Maud", 17)

	fmt.Println(*p1)
	fmt.Println(*p2)

}
