package main

import (
  "fmt"
)

func main() {
  // IF ELSE DECISION
  x := 1
  y := 3

  if (x < y){
    fmt.Println("masuk di kondisi")
  } else {
    fmt.Println("masuk di lainnya")
  }


  //LOOPING SWITCH CASE DECISION 1
  for i := 0; i <= 5; i++ {
    switch i {
    case 1:
      fmt.Println("masuk di kondisi 1")
      break
    case 2:
      fmt.Println("masuk di kondisi 2")
      break
    }
  }

//SWITCH CASE DECISION 2
  for i := 0; i < 10; i++ {
    switch {
    case i < 5:
      fmt.Println("masuk di kondisi -", i)
      break
    case i > 5:
      fmt.Println("masuk di kondisi -", i)
      break
    }
  }

//LOOPING 2
  var j int
  k := 10

  for j < k {
    fmt.Printf("looping ke - %d\n", j)
    j++
  }

  //LOOPING 3,
  //didalam bahasa go tidak ada while maka bisa diatasi dengan looping seperti ini
    l := 0

    for {
      l++

      if(l == 5){
        continue//melewati loop ke -5 dan melanjutkan ke loop berikutnya
      }

      if(l > 10){
        break//berhenti loop ketika nilai l lebih dari 10
      }

      fmt.Printf("hello go ke - %d\n", l)

    }
}
