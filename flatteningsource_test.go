package piper

import (
	"reflect"
	"testing"
)

func TestFlatteningSourceReturnsElementsFromUnderlyingSourcesOfSources(t *testing.T) {

	s := flatteningSource{
		sourceStack: []Source{
			&indexableSource{indexable: reflect.ValueOf([]Source{
				&indexableSource{indexable: reflect.ValueOf([]string{"a", "b"})},
				&indexableSource{indexable: reflect.ValueOf([]string{"c", "d"})},
				&indexableSource{indexable: reflect.ValueOf([]Source{
					&indexableSource{indexable: reflect.ValueOf([]string{"e", "f"})},
				})},
			})},
		},
	}

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

	if !ok {
		t.Fatal("Expected a fourth element but none come next")
	}
	if result[0].String() != "d" {
		t.Fatalf("Expected element 'd' but got %v", result)
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a fifth element but none come next")
	}
	if result[0].String() != "e" {
		t.Fatalf("Expected element 'e' but got %v", result)
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a sixth element but none come next")
	}
	if result[0].String() != "f" {
		t.Fatalf("Expected element 'f' but got %v", result)
	}

	result, ok = s.Next()

	if ok {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
