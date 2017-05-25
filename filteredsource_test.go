package piper

import (
	"reflect"
	"testing"
)

func TestFilteredSourceExcludesFilteredElements(t *testing.T) {
	s := filteredSource{
		source: &indexableSource{indexable: reflect.ValueOf([]string{"a", "b", "c"})},
		filter: reflect.ValueOf(func(x string) bool { return x == "b" }),
	}

	result, ok := s.Next()

	if !ok {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "b" {
		t.Fatalf("Expected element 'b' but got %v", result)
	}

	result, ok = s.Next()

	if ok {
		t.Fatal("Expected no more elements but got %v", result)
	}
}

func TestFilteredSourceUsesRuntimeTypeOfElement(t *testing.T) {
	s := filteredSource{
		source: &indexableSource{indexable: reflect.ValueOf([]interface{}{"a", "b", "c"})},
		filter: reflect.ValueOf(func(x string) bool { return x == "b" }),
	}

	result, ok := s.Next()

	if !ok {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "b" {
		t.Fatalf("Expected element 'b' but got %v", result[0])
	}

	result, ok = s.Next()

	if ok {
		t.Fatal("Expected no more elements but got %v", result)
	}
}
