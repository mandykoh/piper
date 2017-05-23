package piper

import (
	"reflect"
	"strings"
	"testing"
)

func TestProjectedSourceReturnsTransformedElements(t *testing.T) {
	s := projectedSource{
		source:     &indexableSource{indexable: reflect.ValueOf([]string{"a", "b", "c"})},
		projection: reflect.ValueOf(func(x string) string { return strings.ToUpper(x) }),
	}

	result, ok := s.Next()

	if !ok {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "A" {
		t.Fatalf("Expected element 'A' but got %v", result)
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "B" {
		t.Fatalf("Expected element 'B' but got %v", result)
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "C" {
		t.Fatalf("Expected element 'C' but got %v", result)
	}

	result, ok = s.Next()

	if ok {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}

func TestProjectedSourceSupportsMultipleReturnValues(t *testing.T) {
	s := projectedSource{
		source:     &indexableSource{indexable: reflect.ValueOf([]string{"a", "b", "c"})},
		projection: reflect.ValueOf(func(x string) (string, string) { return strings.ToUpper(x), x + "X" }),
	}

	result, ok := s.Next()

	if !ok {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "A" {
		t.Fatalf("Expected element 'A' but got %v", result)
	}
	if result[1].String() != "aX" {
		t.Fatalf("Expected element 'aX' but got %v", result)
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "B" {
		t.Fatalf("Expected element 'B' but got %v", result)
	}
	if result[1].String() != "bX" {
		t.Fatalf("Expected element 'bX' but got %v", result)
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "C" {
		t.Fatalf("Expected element 'C' but got %v", result)
	}
	if result[1].String() != "cX" {
		t.Fatalf("Expected element 'cX' but got %v", result)
	}

	result, ok = s.Next()

	if ok {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
