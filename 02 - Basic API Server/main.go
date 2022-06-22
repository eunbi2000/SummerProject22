package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllBooks)
	myRouter.HandleFunc("/book/{id}", returnSingleBook)
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	Books = []Book{
		{Id: "1", Title: "Hello", Author: "Article Author"},
		{Id: "2", Title: "Hello 2", Author: "Article Author"},
	}
	handleRequests()
}

type Book struct {
	Id     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

var Books []Book

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, book := range Books {
		if book.Id == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}
