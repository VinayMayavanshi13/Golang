package main

import "fmt"

func main() {
	fmt.Println("Welcome to tutorial on If-Else in Golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Something else"
	}

	fmt.Println("Result is:", result)

	/*
		Note: Take care of {} position.
		In Golang you strictly need to use following syntax:
		if condition {

		}

		The lexer won't identify following syntax:
		if condition
		{

		}

	*/

	// var is initialized and used directly
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
}
