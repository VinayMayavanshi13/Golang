package main

import (
	"encoding/json"
	"fmt"
)

/*
Always keep in mind that when you use first letter as small case then you won't want it public
e.g struct course declared below
*/

/*
In Struct instead of Name it is better if we used the term coursename in json
we can do these by writing it in back ticks.

Also in JSON everything is recommended to be in small case.

And if this needs to be used as an API then you won't password to be shown
to do this add this 'json:"-"' besides password in struct

And finally for tags there might be scenario where there is no tag to be mentioned thus in that case use omitempty
this will ensure that null tag does not appear

These method which we used to modify things in Json is called aliases
*/
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("We are learning about JSON")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Java Bootcamp", 599, "LearnCodeOnline.in", "def456", []string{"dsa", "web-dev"}},
		{"Kotlin Bootcamp", 599, "LearnCodeOnline.in", "ghi789", []string{"android", "app-dev"}},
		{"Golang Bootcamp", 399, "LearnCodeOnline.in", "goo321", nil},
	}

	//package this data as JSON data

	// finalJson, err := json.Marshal(myCourses)
	// above we passes in an interface and an interface is a word being borrowed and is an alternative version for mycourses

	//using only marshall provides output in cluttered manner thus we used marshallindent

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	/*
			func json.MarshalIndent(v any, prefix string, indent string) ([]byte, error)
		    MarshalIndent is like Marshal but applies Indent to format the output.
			Each JSON element in the output will begin on a new line beginning with prefix
			followed by one or more copies of indent according to the indentation nesting.
	*/

	//explore the usecases of prefix string

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
   }
	`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
