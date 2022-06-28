package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	MBTI  string `json:"mbti"`
}

var Users []User

func getUsers() []User {
	return Users
}

func returnAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	result, err := json.Marshal(getUsers())
	if err != nil {
		fmt.Print(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	w.WriteHeader(http.StatusOK)
}

func getSingleUser(id int) (*User, error) {
	for _, user := range Users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("no such user exists")
}

func returnSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err.Error())
	}

	user, err := getSingleUser(id)
	if err != nil {
		fmt.Print(err.Error())
	}

	result, err := json.Marshal(user)
	if err != nil {
		fmt.Print(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	w.WriteHeader(http.StatusOK)
}

func createNewUser(user User) {
	Users = append(Users, user)
}

func createNewUserHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	createNewUser(user)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}

func deleteUser(id int) {
	for index, user := range Users {
		if user.Id == id {
			Users = append(Users[:index], Users[index+1:]...)
		}
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err.Error())
	}
	deleteUser(id)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("success"))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the MBTI Test Home Page!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsersHandler).Methods("GET")
	myRouter.HandleFunc("/user", createNewUserHandler).Methods("POST")
	myRouter.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", returnSingleUserHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	Users = []User{}
	handleRequests()
}
