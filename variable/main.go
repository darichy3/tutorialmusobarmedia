package main

import(
  "fmt"
)

func main()  {
  //static type declaration
  var x int
  x = 10

  var y float64
  y = 5.5

  fmt.Println(x)
  fmt.Println(y)

  fmt.Printf("x memiliki tipe data: %T\n", x)
  fmt.Printf("y memiliki tipe data: %T\n", y)

  //dynamic type declaration
  a := "fedora"
  b := 28;
  fmt.Println("version ", a , b)

}
