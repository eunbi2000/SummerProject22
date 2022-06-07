package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

//increment counter 진짜 ++ 하는지 and handler response 제대로?
func TestIncrement(t *testing.T) {
	for j := 1; j <= 10; j++ {
		compare := j
		counter := Increment(j - 1)
		if counter != compare {
			t.Errorf("expected '%d' but got '%d'", compare, counter)
		}
	}
}
func TestIncrementCounterHandler(t *testing.T) {
	for j := 1; j <= 10; j++ {
		expected := strconv.Itoa(j)
		req := httptest.NewRequest(http.MethodGet, "/increment?", nil)
		w := httptest.NewRecorder()
		IncrementCounterHandler(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if string(data) != expected {
			t.Errorf("Expected %v but got %v", expected, string(data))
		}
	}
}
