package main

import (
	"reflect"
	"testing"
)

func TestBasicCounter(t *testing.T) {
	compare := 0
	if reflect.TypeOf(counter) != reflect.TypeOf(compare) {
		t.Errorf("expected '%d' but got '%d'", reflect.TypeOf(compare), reflect.TypeOf(counter))
	}
}
