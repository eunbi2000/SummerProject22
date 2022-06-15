package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var tests = []struct {
	x   int
	exp int
}{
	{0, 1},
	{1, 2},
	{2, 3},
	{3, 4},
	{4, 5},
	{5, 6},
	{6, 5}, // incorrect test case
}

func TestIncrement(t *testing.T) {
	for _, test := range tests {
		compare := increment(test.x)
		if compare != test.exp {
			t.Errorf("expected '%d' but got '%d'", compare, test.exp)
		}
	}
}
func TestIncrementCounterHandler(t *testing.T) {
	for _, test := range tests {
		expected := strconv.Itoa(test.exp)
		req := httptest.NewRequest(http.MethodGet, "/increment", nil)
		w := httptest.NewRecorder()
		incrementCounterHandler(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		result := string(data)
		if result != expected {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	}
}
