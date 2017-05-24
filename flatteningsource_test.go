package piper

import (
	"reflect"
	"testing"
)

func TestFlatteningSourceFormsCartesianProductOfMultipleReturnValues(t *testing.T) {

	s := flatteningSource{
		source: projectedSource{
			source: &indexableSource{indexable: reflect.ValueOf([]string{"dummy"})},
			projection: reflect.ValueOf(func(v string) ([]string, string, []string) {
				return []string{"1", "2"}, "x", []string{"a", "b"}
			}),
		},
	}

	result, ok := s.Next()

	if !ok {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "1" {
		t.Errorf("Expected element '1' but got %v", result[0])
	}
	if result[1].String() != "x" {
		t.Errorf("Expected element 'x' but got %v", result[1])
	}
	if result[2].String() != "a" {
		t.Errorf("Expected element 'a' but got %v", result[2])
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "1" {
		t.Errorf("Expected element '1' but got %v", result[0])
	}
	if result[1].String() != "x" {
		t.Errorf("Expected element 'x' but got %v", result[1])
	}
	if result[2].String() != "b" {
		t.Errorf("Expected element 'b' but got %v", result[2])
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "2" {
		t.Errorf("Expected element '2' but got %v", result[0])
	}
	if result[1].String() != "x" {
		t.Errorf("Expected element 'x' but got %v", result[1])
	}
	if result[2].String() != "a" {
		t.Errorf("Expected element 'a' but got %v", result[2])
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a fourth element but none come next")
	}
	if result[0].String() != "2" {
		t.Errorf("Expected element '2' but got %v", result[0])
	}
	if result[1].String() != "x" {
		t.Errorf("Expected element 'x' but got %v", result[1])
	}
	if result[2].String() != "b" {
		t.Errorf("Expected element 'b' but got %v", result[2])
	}

	result, ok = s.Next()

	if ok {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}

func TestFlatteningSourceUnwrapsSlicesFromUnderlyingSource(t *testing.T) {

	s := flatteningSource{
		source: &indexableSource{indexable: reflect.ValueOf([][]string{
			[]string{"a", "b"},
			[]string{"c", "d"},
		})},
	}

	result, ok := s.Next()

	if !ok {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "a" {
		t.Errorf("Expected element 'a' but got %v", result[0])
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "b" {
		t.Errorf("Expected element 'b' but got %v", result[0])
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "c" {
		t.Errorf("Expected element 'c' but got %v", result[0])
	}

	result, ok = s.Next()

	if !ok {
		t.Fatal("Expected a fourth element but none come next")
	}
	if result[0].String() != "d" {
		t.Errorf("Expected element 'd' but got %v", result[0])
	}

	result, ok = s.Next()

	if ok {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
