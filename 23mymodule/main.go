package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Mod in GoLang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

/*
It's completely fine to run the code with using go mod init.
in this case let us do : go mod init github.com/VinayMayavanshi13/mymodules
*/

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang Series"))
}

/*
In our mod file we have go version 1.18
Command to change the go version :
go mod edit -go 1.16 => This changes the version to 1.16
To change the module :
go mod edit -module 1.16

Another command to look for is :
go mod vendor
*/
