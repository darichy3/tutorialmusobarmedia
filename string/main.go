package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var myString = "Hello Golang"
	var myStringTwo = "200"

	fmt.Println(myString)
	fmt.Println(len(myString))

	myStringUpper := strings.ToUpper(myString)
	fmt.Println(myStringUpper)

	myStringLower := strings.ToLower(myString)
	fmt.Println(myStringLower)

	resultContain := strings.Contains(myString, "Go")
	fmt.Println(resultContain)

	resultSplit := strings.Split(myString, " ") //return dinamyc array / slice
	for _, v := range resultSplit {
		fmt.Println(v)
	}

	resultConvInt, err := strconv.Atoi(myStringTwo)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultConvInt * 4)

	myStrReplace := strings.Replace(myString, "Golang", "Java", 1)
	fmt.Println(myStrReplace)

	resultConvStr := strconv.Itoa(10)

	fmt.Println(resultConvStr)
}
