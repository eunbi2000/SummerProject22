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

func getSingleUser(id int) *User {
	for _, user := range Users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func createNewUser(user User) {
	Users = append(Users, user)
}

func deleteUser(id int) (exists bool) {
	for index, user := range Users {
		if user.Id == id {
			Users = append(Users[:index], Users[index+1:]...)
			return true
		}
	}
	return false
}

func returnAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	result, err := json.Marshal(getUsers())
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	w.WriteHeader(http.StatusOK)
}

func returnSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	user := getSingleUser(id)
	if user != nil {
		result, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("no such user"))
		w.WriteHeader(http.StatusOK)
	}

}

func createNewUserHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	createNewUser(user)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("delete unsuccessful"))
	}
	success := deleteUser(id)
	if success {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("success"))
	} else {
		log.Fatal(errors.New("user doesn't exist"))
	}
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
