package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	// var ptr *int
	// fmt.Println("Value of Pointer is ", ptr)

	myNumber := 24

	var ptr = &myNumber //here ptr is not a pointer
	fmt.Println("Value of actual Pointer is ", ptr)
	fmt.Println("Value of actual Pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is:", myNumber)

}
