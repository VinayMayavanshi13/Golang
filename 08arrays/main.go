package main

import "fmt"

func main() {

	fmt.Println("Welcome to Arrays in Golang")

	var fruitlist [4]string

	fruitlist[0] = "Mango"
	fruitlist[1] = "Papaya"
	//fruitlist[2] = "Kiwi"
	fruitlist[3] = "Banana"

	fmt.Println("Fruit list is: ", fruitlist)

	//Output = Fruit list is:  [Mango Papaya  Banana]
	/*
		Here Space b/w Mango and Papaya is 1 whereas the same b/w Papaya and Banana is 2
		since we have not declared anything at 2nd index in the array
	*/

	fmt.Println("Length of Fruit list is: ", len(fruitlist))

	//Output : Length of Fruit list is:  4
	/*
		Its an declared value.
	*/

	var vegList = [3]string{"Brocolli", "Mushroom", "Beans"}

	fmt.Println("Vegetable list is: ", vegList)
	fmt.Println("Length of Vegetable list is: ", len(vegList))

}
