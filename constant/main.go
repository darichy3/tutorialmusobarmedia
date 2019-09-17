package main

import (
  "fmt"
)

const A string = "INI ADALAH A GLOBAL CONSTANT"
func main()  {
  fmt.Println(A)

  const X = 10

  fmt.Println("Nilai local constant x =",X)

  y := 200/X

  fmt.Println("Nilai 200 / constant X =",y)

}
