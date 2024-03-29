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

func increment(c int) (counter int) {
	c++
	counter = c
	return
}

func incrementCounterHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, strconv.Itoa(increment(counter)))
	mutex.Unlock()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/increment", incrementCounterHandler)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
