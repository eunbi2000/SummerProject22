package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the MBTI Test Home Page!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsers)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", returnSingleUser)
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	Users = []User{
		{Id: "1", Name: "John Smith", Email: "example1@gmail.com", MBTI: "INTP"},
		{Id: "2", Name: "Jane Doe", Email: "example2@gmail.com", MBTI: "ENFP"},
	}
	handleRequests()
}

type User struct {
	Id    string `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
	MBTI  string `json:"MBTI"`
}

var Users []User

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	prettyprint(w, Users)
}

func prettyprint(w http.ResponseWriter, data interface{}) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(data)
	return
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, user := range Users {
		if user.Id == key {
			prettyprint(w, user)
		}
	}
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	Users = append(Users, user)

	prettyprint(w, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, user := range Users {
		if user.Id == id {
			Users = append(Users[:index], Users[index+1:]...)
		}
	}
}
