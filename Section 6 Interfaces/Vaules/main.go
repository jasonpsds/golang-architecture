package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("From a person ", p.first)
}

// interface
type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	fmt.Printf("%T\n", p1)

	// a value can be of more than one type
	// p1 is both type person and human
	var x human
	x = p1
	fmt.Printf("%T\n", x)
}
