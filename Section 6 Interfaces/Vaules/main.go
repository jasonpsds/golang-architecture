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

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I'am a seceret agent, the name is", sa.first)
}

func main() {
	p1 := person{
		first: "Miss Moneypenny",
	}

	sa1 := secretAgent{
		person: person{
			first: "Larry Flint",
		},
		ltk: true,
	}

	fmt.Printf("%T\n", p1)

	// a value can be of more than one type
	// p1 is both type person and human
	var x, y human
	x = p1
	y = sa1
	x.speak()
	y.speak()
	// fmt.Printf("%T\n", x)
}
