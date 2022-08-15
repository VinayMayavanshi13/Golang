package main

import "fmt"

//When we write first letter of variable as Capital it means it is a public variable same as declaring with public keyword.
const LoginToken string = "ASFET4589" //Public

func main() {
	var username string = "vinay"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45546789568789553135
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var smallFloat2 float64 = 255.45546789568789553135
	fmt.Println(smallFloat2)
	fmt.Printf("Variable is of type: %T \n", smallFloat2)

	// default values and aliases
	var intVal int // here default value of int will be 0.
	fmt.Println(intVal)
	fmt.Printf("Variable is of type: %T \n", intVal)

	// implicit type

	var website = "www.github.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	//This also Works fine
	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)

	//This no var style declaration can't work for global scope

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
