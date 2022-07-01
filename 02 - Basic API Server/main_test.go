package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestReturnAllUsersHandler(t *testing.T) {
	expected := []User{
		{Id: 1, Name: "John Smith", Email: "example1@gmail.com", MBTI: "INTP"},
		{Id: 2, Name: "Jane Doe", Email: "example2@gmail.com", MBTI: "ENFP"},
	}
	Users = append(expected)
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnAllUsersHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	exp, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}
	if rr.Body.String() != string(exp) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body, string(exp))
	}
}

func TestReturnSingleUserHandler(t *testing.T) { //ask how to set specific url param
	expected := User{
		Id: 1, Name: "John Smith", Email: "example1@gmail.com", MBTI: "INTP",
	}
	Users = append(Users, expected)
	req, err := http.NewRequest("GET", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", returnSingleUserHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	exp, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	want := getSingleUser(1)
	if want == nil {
		t.Errorf(err.Error())
	}

	if !(reflect.DeepEqual(*want, expected)) {
		t.Errorf("Didn't get correct user, got %v want %v", want, expected)
	}
	if rr.Body.String() != string(exp) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body, string(exp))
	}
}

func TestCreateNewUserHandler(t *testing.T) { //create 했을때 getalluser 통해 확인할수있는지 아니면 func자체에 넣는지
	inputUser := User{Id: 3, Name: "Test", Email: "test@gmail.com", MBTI: "ISFJ"}

	result, err := json.Marshal(inputUser)
	if err != nil {
		t.Fatal(err.Error())
	}

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(result))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createNewUserHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := &User{Id: 3, Name: "Test", Email: "test@gmail.com", MBTI: "ISFJ"}

	want := getSingleUser(3)
	if want == nil {
		t.Errorf(err.Error())
	}

	if !(reflect.DeepEqual(want, expected)) {
		t.Errorf("Didn't create user, got %v want %v", want, expected)
	}
	if rr.Body.String() != "success" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body, "success")
	}
}

func TestDeleteUserHandler(t *testing.T) { //ask how to set specific url param
	Users = []User{
		{Id: 1, Name: "John Smith", Email: "example1@gmail.com", MBTI: "INTP"},
		{Id: 2, Name: "Jane Doe", Email: "example2@gmail.com", MBTI: "ENFP"},
	}
	req, err := http.NewRequest("DELETE", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}

	want := getSingleUser(1)
	// want no such user error
	if want != nil {
		t.Errorf(err.Error())
	}

	if rr.Body.String() != "success" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body, "success")
	}
}
