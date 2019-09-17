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
}
