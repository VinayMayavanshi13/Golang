package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjkkn546lkj"

// In urls,the & i.e ampersand sign acts as comma.

func main() {
	fmt.Println("Welcome to Url Handlings")
	fmt.Println(myurl)

	//parsing

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) // type came out to be url.Values means a key - value pair

	fmt.Println(qparams["coursename"])

	// for _, val := range qparams {
	// 	fmt.Println("Param is: ", val)
	// }

	for index, val := range qparams {
		fmt.Printf("%v is: %v\n", index, val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=vinay",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Printf(anotherUrl)

}
