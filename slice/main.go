package main

//dinamis array, karena slice bisa bertambah panjangnya secara otomatis tidak seperti array

import (
	"fmt"
)

func main() {
	mySlice := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}

	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	mySliceStr := []string{"Zen", "Rizk", "Dwi", "Bry"}
	mySliceStr = append(mySliceStr, "Test", "Yups")

	for _, v := range mySliceStr {
		fmt.Println(v)
	}
}
