package main

import (
  "fmt"
)

/*
  This is about function return function
  The other comment in golang
*/



func main()  {
  nextValue := getNumber()

  fmt.Println(nextValue())
  fmt.Println(nextValue())
  fmt.Println(nextValue())

  lv := love("Rizki")
  fmt.Println(lv("golang"))
  fmt.Println(lv("java"))
  fmt.Println(lv("php"))
}

func getNumber() func() int  {
  x:= 0

  return func () int {
    x = x + 1
    return x
  }
}

func love(nama string) func(string) string  {
  return func (sesuatu string) string  {
    return fmt.Sprintf("%s love %s", nama, sesuatu)
  }
}
