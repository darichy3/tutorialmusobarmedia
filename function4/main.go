package main

import (
  "fmt"
)
/*
  function as param
*/

func main()  {

  f := func(v string) bool  {
    return v == "Golang"
  }

  fmt.Println(match("Golang", f))

}

func match(v string, f func(string) bool) bool  {
  if f(v) {
    return true
  }
  return false
}
