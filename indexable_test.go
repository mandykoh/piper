package piper

import (
	"reflect"
	"testing"
)

func TestIndexableSourceReturnsElementsFromSlice(t *testing.T) {
	indexable := &indexable{items: reflect.ValueOf([...]string{"a", "b", "c"})}

	var s Source = indexable.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "a" {
		t.Fatalf("Expected element 'a' but got %v", result)
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "b" {
		t.Fatalf("Expected element 'b' but got %v", result)
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "c" {
		t.Fatalf("Expected element 'c' but got %v", result)
	}

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
