package piper

import (
	"reflect"
	"testing"
)

func TestIndexableSourceReturnsElementsFromSlice(t *testing.T) {
	s := indexableSource{indexable: reflect.ValueOf([...]string{"a", "b", "c"})}

	result, ok := s.Next()

	if !ok {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "a" {
		t.Fatalf("Expected element 'a' but got %v", result)
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "b" {
		t.Fatalf("Expected element 'b' but got %v", result)
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "c" {
		t.Fatalf("Expected element 'c' but got %v", result)
	}

	result, ok = s.Next()

	if ok {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
