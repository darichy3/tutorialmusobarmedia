package main

import (
	"fmt"
)

func main() {
	p := &Person{
		ID:   1,
		Name: "Joko",
		Age:  27,
	}

	p.changeName("Rizki")

	fmt.Println(p.GetID())
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
}

type Person struct {
	ID   int
	Name string
	Age  int
}

//function ini merupakan method, karena fuction ini merupakan member dari Person
func (p *Person) GetID() int { //(p *Person) ini merupakan receiver
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) changeName(newName string) { //method merubah nama, dimana konsep sama kayak setter
	p.Name = newName
}
