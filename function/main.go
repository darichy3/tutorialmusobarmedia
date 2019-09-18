package main

import (
  "fmt"
)

func main()  {
  fmt.Println("fungsi main")

  x := 2
  y := 6
  z := add(x,y)

  fmt.Println(x)
  fmt.Println(y)
  fmt.Println("niali x + y =",z)

  nama := "Rizki"
  fmt.Println(salam(nama))

  a := "Rizki"
  b := 1
  fmt.Println(swap(a,b))

}

func add( x int, y int) int {
  a := x + y
  return a
}

func salam(nama string) string{
  result := fmt.Sprintf("Hai %s !!", nama)
  return result
}

func swap(a string, b int) (int, string){
  return b, a
}
