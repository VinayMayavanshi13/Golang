package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods")
	vinay := User{"Vinay", "vinay@own.com", true, 22}

	// fmt.Println(vinay)
	// fmt.Printf("Vinay details are: %+v\n", vinay) // note the use of %+v for struct datatype
	// fmt.Printf("Name is %v and email is %v.", vinay.Name, vinay.Email)

	vinay.GetStatus()
	vinay.NewMail()
	fmt.Printf("Name is %v and email is %v.", vinay.Name, vinay.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

/*
If in the struct i declared a new var say oneAge int then that won't be exportable since first letter of that
variable is in smallcase.
*/

// method :

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
