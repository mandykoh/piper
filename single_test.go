package piper

import (
	"reflect"
	"testing"
)

func TestSingleReturnsValuesForOneItem(t *testing.T) {
	single := &single{values: []interface{}{"a", 1}}

	var s WrappedSource = single.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "a" {
		t.Fatalf("Expected element 'a' but got %v", result[0])
	}
	if result[1].Int() != 1 {
		t.Fatalf("Expected element 1 but got %v", result[1])
	}

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
