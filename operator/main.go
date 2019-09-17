package main

import (
  "fmt"
)

func main()  {
  x := 10
  y := 5

  //aritmatic operator
  fmt.Println("x =",x)
  fmt.Println("y =",y)

  //penambahan / additional
  zAdd := x + y
  fmt.Println("x + y =",zAdd)

  //pengurangan / substracts
  zSub := x - y
  fmt.Println("x - y =",zSub)

  //perkalian / multiplies
  zMul := x * y
  fmt.Println("x * y =",zMul)

  //pembagian / divides
  zDiv := x / y
  fmt.Println("x / y =",zDiv)

  //modulus
  zMod := x % y
  fmt.Println("x % y =",zMod)
  fmt.Println("x / y =",zDiv)

  //relational operator
  i := 4
  j := 2
  fmt.Println("i =",i)
  fmt.Println("j =",j)

  fmt.Println("i == j =",i == j) //equals
  fmt.Println("i != j =",i != j) //not equalse
  fmt.Println("i > j =",i > j) //grether than
  fmt.Println("i < j =",i < j) //less then
  fmt.Println("i >= j =",i >= j) //grether than
  fmt.Println("i <= j =",i <= j) //less then
}
