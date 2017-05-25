package piper

import (
	"reflect"
	"testing"
)

func TestFilterExcludesFilteredElements(t *testing.T) {
	indexable := &indexable{items: reflect.ValueOf([]string{"a", "b", "c"})}
	f := filter{
		source: indexable.Source,
		test:   reflect.ValueOf(func(x string) bool { return x == "b" }),
	}

	var s Source = f.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "b" {
		t.Fatalf("Expected element 'b' but got %v", result)
	}

	result, s = s()

	if s != nil {
		t.Fatal("Expected no more elements but got %v", result)
	}
}

func TestFilteredSourceUsesRuntimeTypeOfElement(t *testing.T) {
	indexable := &indexable{items: reflect.ValueOf([]interface{}{"a", "b", "c"})}
	f := filter{
		source: indexable.Source,
		test:   reflect.ValueOf(func(x string) bool { return x == "b" }),
	}

	var s Source = f.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "b" {
		t.Fatalf("Expected element 'b' but got %v", result[0])
	}

	result, s = s()

	if s != nil {
		t.Fatal("Expected no more elements but got %v", result)
	}
}
