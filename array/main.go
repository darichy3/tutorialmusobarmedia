package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arrStr := [5]string{"Rizki", "Dwi", "Go", "Lang", "Learn"}

	fmt.Println(arrStr[4])

	//array multi dimensi 2 baris dan 3 kolom, artinya ada 2 baris dan 3 kolom
	arrMulti := [2][3]string{{"A", "B", "C"}, {"D", "E", "F"}}
	//mengambil nilai dari array multi dimensi baris ke 1, dan kolom ke 0 yaitu "D" karena arrai itu di mulai dari 0
	fmt.Println(arrMulti[1][0])
}
