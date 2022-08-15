package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")

	//There is no inheritance in Go Lang; No Super or Parent

	vinay := User{"Vinay", "vinay@own.com", true, 22}

	fmt.Println(vinay)
	fmt.Printf("Vinay details are: %+v\n", vinay) // note the use of %+v for struct datatype
	fmt.Printf("Name is %v and email is %v.", vinay.Name, vinay.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
