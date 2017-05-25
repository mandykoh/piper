package piper

import (
	"reflect"
	"testing"
)

func TestFlattenFormsCartesianProductOfMultipleReturnValues(t *testing.T) {
	many := &many{items: reflect.ValueOf([]string{"dummy"})}
	projector := projector{
		source: many.Source,
		projections: []reflect.Value{
			reflect.ValueOf(func(v string) ([]string, string, []string) {
				return []string{"1", "2"}, "x", []string{"a", "b"}
			}),
		},
	}
	flattener := &flatten{source: projector.Source}

	var s Source = flattener.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
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

	result, s = s()

	if s == nil {
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

	result, s = s()

	if s == nil {
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

	result, s = s()

	if s == nil {
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

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}

func TestFlattenUnwrapsSlicesFromUnderlyingSource(t *testing.T) {
	many := &many{items: reflect.ValueOf([][]string{
		[]string{"a", "b"},
		[]string{"c", "d"},
	})}
	flattener := &flatten{source: many.Source}

	var s Source = flattener.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "a" {
		t.Errorf("Expected element 'a' but got %v", result[0])
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "b" {
		t.Errorf("Expected element 'b' but got %v", result[0])
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "c" {
		t.Errorf("Expected element 'c' but got %v", result[0])
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a fourth element but none come next")
	}
	if result[0].String() != "d" {
		t.Errorf("Expected element 'd' but got %v", result[0])
	}

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}

func TestFlattenUsesRuntimeTypeToDetermineIndexables(t *testing.T) {
	many := &many{items: reflect.ValueOf([][]string{
		[]string{"a", "b"},
		[]string{"c", "d"},
	})}
	projector := projector{
		source:      many.Source,
		projections: []reflect.Value{reflect.ValueOf(func(x []string) interface{} { return x })},
	}
	flattener := &flatten{source: projector.Source}

	var s Source = flattener.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "a" {
		t.Errorf("Expected element 'a' but got %v", result[0])
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "b" {
		t.Errorf("Expected element 'b' but got %v", result[0])
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "c" {
		t.Errorf("Expected element 'c' but got %v", result[0])
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a fourth element but none come next")
	}
	if result[0].String() != "d" {
		t.Errorf("Expected element 'd' but got %v", result[0])
	}

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
