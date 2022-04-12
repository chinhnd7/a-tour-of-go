package main

import "fmt"

func main() {
	array := [4]int{1, 2, 3, 4}
	var pointer *[4]int

	pointer = &array
	fmt.Println(*pointer)

	stringTest := "Love the way you lie"

	var pointerStr *string
	pointerStr = &stringTest
	fmt.Println(pointerStr)

}
