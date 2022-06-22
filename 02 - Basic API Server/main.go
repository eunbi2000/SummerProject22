package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/books", returnAllBooks)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	Books = []Book{
		{Title: "Hello", Author: "Article Author"},
		{Title: "Hello 2", Author: "Article Author"},
	}
	handleRequests()
}

type Book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

var Books []Book

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Books)
}
