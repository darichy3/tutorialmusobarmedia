package main

import (
	"fmt"
)

func main() {
	var p Person
	p.ID = 1
	p.Name = "Rizki"
	p.Age = 29

	q := Person{
		ID:   2,
		Name: "Delta",
		Age:  24,
	}

	printPerson(p)
	printPerson(q)
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func printPerson(p Person) { //struct sebagai parameter function
	fmt.Println("ID =", p.ID)
	fmt.Println("Name =", p.Name)
	fmt.Println("Age =", p.Age)
}
