package main

import "fmt"

type person struct {
	first string
}

// Faux db storage system
type mongo map[int]person
type postg map[int]person

// methods
func (m mongo) save(n int, p person) {
	m[n] = p
}

func (m mongo) receive(n int) person {
	return m[n]
}

func (pg postg) save(n int, p person) {
	pg[n] = p
}

func (pg postg) receive(n int) person {
	return pg[n]
}

// accessor interface
type accessor interface {
	save(n int, p person)
	receive(n int) person
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
	return a.receive(n)
}

func main() {
	// create database
	dbm := mongo{}
	dbp := postg{}

	// some data
	p1 := person{
		first: "Bob",
	}
	p2 := person{
		first: "Jacinda",
	}

	// store data in the db
	// mongo
	put(dbm, 1, p1)
	put(dbm, 2, p2)

	// pastg
	put(dbp, 1, p1)
	put(dbp, 2, p2)

	// get some data from the db and output to console.
	// mongo
	fmt.Println(get(dbm, 1))
	fmt.Println(get(dbm, 2))

	// postg
	fmt.Println(get(dbp, 1))
	fmt.Println(get(dbp, 2))
}
