package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in Golang")

	content := "This has to go in the file."

	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// }

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./mygofile.txt")

	/*
			readFile Output :
							 Text data inside the file is
							 [84 104 105 115 32 104 97 115 32 116 111 32 103 111 32 105 110 32 116 104 101 32 102 105 108 101 46]

		    As we can see here, the data is read in bytes and not as plain string.

			but we can parse this by passing it in string() :
				fmt.Println("Text data inside the file in string format: \n", string(databyte))

				Output :
							           Text data inside the file in string format:
				 				       This has to go in the file.

	*/
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file in string format: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
