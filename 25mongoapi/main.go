package main

import (
	"fmt"
	"log"
	"net/http"

	"mongoapi/router"
)

func main() {
	fmt.Println("Welcome to MongoAPI")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
