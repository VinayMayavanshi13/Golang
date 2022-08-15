package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RU"] = "Rust"

	fmt.Println("List of all Langauges: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RU")

	fmt.Println("List of all Langauges: ", languages)

	//Loops are interesting in GoLang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)

		// fmt.Printf("%v shorts for %v\n", key, value)
	}

	for _, value := range languages {

		fmt.Printf("For key v, value is %v\n", value)

	}
}
