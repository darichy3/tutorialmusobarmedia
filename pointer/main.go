package main

import (
	"fmt"
)

func main() {

	var hello string = "Hello"
	var strPtr *string //tanda * sebagai penanda merujuk ke alamat memori

	strPtr = &hello //tanda amp / & sebagai pengakses ke alamat memori

	fmt.Println(&hello)
	fmt.Println(strPtr)

	fmt.Println(hello)
	change(hello)
	fmt.Println(hello)
	changePtr(&hello)
	fmt.Println(hello)

}

//copy by value
func change(v string) {
	v = v + " Golang"
}

//copy by reference
func changePtr(v *string) {
	*v = *v + " Golang"
}
