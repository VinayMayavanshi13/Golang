package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to this Program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("On a Scale of 1 - 10 enter your excitement level of learning Go:")
	//this we have used println so no need to use new line character

	// comma ok syntax or err ok syntax for catching errors or so

	//because in Go lang paradigm we do not have try and catch error the error are treated as variables having true - false values or
	//something else.

	//either you will have input or error,that error can be hold up in a variable or just use an underscore

	//below instead of underscore we can have variable name err.
	//it is like try and catch syntax only, input being try and _ being catch.
	input, _ := reader.ReadString('\n')
	//here you are reading string till a new line charcter is encountered.

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input) // suprisingly it gives type as String

}

// for user input read bufio package on Go documentation.
