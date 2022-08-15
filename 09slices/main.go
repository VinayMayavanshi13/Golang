package main

import "fmt"

func main() {
	// fmt.Println("Welcome to Slices")

	// var fruitList = []string{"Kiwi", "Papaya", "Jackfruit"}

	// fmt.Printf("Type of fruitList is %T\n", fruitList)
	// fmt.Println("The fruitList is ", fruitList)
	// fmt.Println("The Length of fruitList is ", len(fruitList))

	// fruitList = append(fruitList, "Mango", "Orange")

	// fmt.Println("The fruitList is ", fruitList)
	// fmt.Println("The Length of fruitList is ", len(fruitList))

	// fruitList = append(fruitList[1:])
	// fmt.Println("Now after slicing 1:, the fruitList is ", fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println("Now after slicing 1:3, the fruitList is ", fruitList)

	// fruitList = append(fruitList[:3])
	// fmt.Println("Now after slicing :3, the fruitList is ", fruitList) // it will start from default i.e 0

	// fruitList = append(fruitList[3:])
	// fmt.Println("Now after slicing 3:, the fruitList is ", fruitList)

	// fmt.Println("The fruitList is ", fruitList)

	/*
		So slicing permanently makes changes to the array.
	*/

	// highScores := make([]int, 4)

	// highScores[0] = 234
	// highScores[1] = 345
	// highScores[2] = 496
	// highScores[3] = 786

	//highScores[4] = 786 => This will throw up an error saying index out of range

	// fmt.Println(highScores)

	// highScores = append(highScores, 555, 987, 367) // this works fine as append reallocates the memory to the slice.
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on the index

	var courses = []string{"ReactJS", "JavaScript", "Kotlin", "Python", "Machine Learning"}

	fmt.Println("The Courses you interested in are: ", courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("After removal process we have: ", courses)
}
