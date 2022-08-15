package main

import "fmt"

/*
Notes:
1. In Golang you are not allowed to write a function within a function.
2. In Golang you can also create a func with no name, but immediately invoke it. e.g :

func (){
	fmt.Println("I am A function with no name.")
}()

As of now this might throw an error, we will discuss this later.

*/

func main() {
	fmt.Println("Welcome to Functions!!!")
	greeter()

	result1 := adder(3, 5)

	fmt.Println("Result1 is: ", result1)

	result2, myMessage := proAdder(1, 2, 3, 4, 5)

	fmt.Println("Result2 is: ", result2)
	fmt.Println("Message is: ", myMessage)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// when you don't know number of arguments do following:
// below we are also returning two values int and string
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hii, I am from Pro func"
}

func greeter() {
	fmt.Println("Hello from Golang")
}
