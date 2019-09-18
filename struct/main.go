package main

import (
	"fmt"
)

type Vector struct {
	x, y int // jika tipe data sama bisa deklarasi dengan cara ini
}

type Player struct {
	ID   int
	Name string
}

func main() {

	var v Vector
	v.x = 10
	v.y = 5

	fmt.Println(v) //mengakses vektor secara keseluruhan

	fmt.Println("X =", v.x)
	fmt.Println("Y =", v.y)

	player1 := Player{ID: 1, Name: "Neymar"}
	player2 := Player{ID: 2, Name: "Ronaldo"}

	fmt.Println(player1.ID)
	fmt.Println(player1.Name)
	fmt.Println(player2.ID)
	fmt.Println(player2.Name)
}
