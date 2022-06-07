package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func Increment(c int) (incremented int) {
	c++
	counter = c
	return counter
}

func IncrementCounterHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	result := Increment(counter)
	fmt.Fprintf(w, strconv.Itoa(result))
	mutex.Unlock()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/increment", IncrementCounterHandler)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
