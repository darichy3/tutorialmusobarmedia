package main

import (
  "fmt"
)

func main()  {
  //function as value (anonimous function)
  add := func (x, y int) int  {
    return x + y
  }

  fmt.Println(add(2, 4))
}
