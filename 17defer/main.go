package main

import "fmt"

// func main() {
// 	fmt.Println("Welcome to Defers in Golang")

// 	defer fmt.Println("Hello")
// 	fmt.Println("World")

// 	/*
// 		Since line 8 has been defered so it will excuted later, first line 9 will be executed thus we will get "World Hello" on console
// 		instead of Hello World
// 	*/

// }

func main() {
	// fmt.Println("Welcome to Defers in Golang")

	// defer fmt.Println("Hello")
	// defer fmt.Println("One")
	// defer fmt.Println("Two")
	// fmt.Println("World")
	/*
	   Output :
	   			Welcome to Defers in Golang
	   			World
	   			Two
	   			One
	   			Hello
	   As we can see here defer followed the principle of LIFO i.e Last In First Out.
	*/

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
