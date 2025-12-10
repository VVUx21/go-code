package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("this is a test")
	fmt.Println("this is a new line")
	greeter()
	r:=mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")
	http.ListenAndServe(":8080",r)
}

func greeter()  {
	fmt.Println("welcome to go modules")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to the Home Page </h1>"))
}

//The ENv variable GOPATH is used to specify the location of your workspace
//By default, it is set to a directory named "go" in your home directory
//When you run go commands, Go will look for packages and modules in the directories specified by GOPATH
//cache directory is used to store downloaded modules
//pkg directory is used to store compiled package files
//src directory is used to store source code files for your Go projects
//go sum file is used to verify that modules used in your project have not been tampered with.
//It contains cryptographic hashes of the module versions specified in your go.mod file.
//go mod vendor command is used to create a vendor directory in your project
//This directory contains copies of all the dependencies required by your project