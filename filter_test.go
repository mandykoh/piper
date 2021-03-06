package piper

import (
	"reflect"
	"testing"
)

func TestFilterExcludesFilteredElements(t *testing.T) {
	many := &many{items: reflect.ValueOf([]string{"a", "b", "c"})}
	f := filter{
		source: many.Source,
		test:   reflect.ValueOf(func(x string) bool { return x == "b" }),
	}

	var s WrappedSource = f.Source
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
		t.Fatalf("Expected no more elements but got %v", result)
	}
}

func TestFilterUsesRuntimeTypeOfElement(t *testing.T) {
	many := &many{items: reflect.ValueOf([]interface{}{"a", "b", "c"})}
	f := filter{
		source: many.Source,
		test:   reflect.ValueOf(func(x string) bool { return x == "b" }),
	}

	var s WrappedSource = f.Source
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
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
