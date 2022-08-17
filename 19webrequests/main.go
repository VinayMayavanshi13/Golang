package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.github.com"

func main() {
	fmt.Println("Welcome to Web requests in Golang.")

	response, err := http.Get(url)

	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content := string(databytes)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
