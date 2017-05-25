package piper

import (
	"reflect"
	"strings"
	"testing"
)

func TestProjectorCombinesValuesReturnedFromMultipleProjections(t *testing.T) {
	many := &many{items: reflect.ValueOf([]interface{}{"a", "b"})}
	projector := &projector{
		source: many.Source,
		projections: []reflect.Value{
			reflect.ValueOf(func(x string) (string, int) { return strings.ToUpper(x), 1 }),
			reflect.ValueOf(func(x string) (string, int) { return x + x, 2 }),
		},
	}

	var s WrappedSource = projector.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if count := len(result); count != 4 {
		t.Fatalf("Expected 4 return values but got %d", count)
	}
	if result[0].String() != "A" {
		t.Fatalf("Expected element 'A' but got %v", result[0])
	}
	if result[1].Int() != 1 {
		t.Fatalf("Expected element 1 but got %v", result[1])
	}
	if result[2].String() != "aa" {
		t.Fatalf("Expected element 'aa' but got %v", result[0])
	}
	if result[3].Int() != 2 {
		t.Fatalf("Expected element 2 but got %v", result[1])
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if count := len(result); count != 4 {
		t.Fatalf("Expected 4 return values but got %d", count)
	}
	if result[0].String() != "B" {
		t.Fatalf("Expected element 'B' but got %v", result[0])
	}
	if result[1].Int() != 1 {
		t.Fatalf("Expected element 1 but got %v", result[1])
	}
	if result[2].String() != "bb" {
		t.Fatalf("Expected element 'bb' but got %v", result[0])
	}
	if result[3].Int() != 2 {
		t.Fatalf("Expected element 2 but got %v", result[1])
	}
}

func TestProjectorPassesArgumentsUsingRuntimeType(t *testing.T) {
	many := &many{items: reflect.ValueOf([]interface{}{"a", "b", "c"})}
	projector := projector{
		source:      many.Source,
		projections: []reflect.Value{reflect.ValueOf(func(x string) string { return strings.ToUpper(x) })},
	}

	var s WrappedSource = projector.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "A" {
		t.Fatalf("Expected element 'A' but got %v", result)
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "B" {
		t.Fatalf("Expected element 'B' but got %v", result)
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "C" {
		t.Fatalf("Expected element 'C' but got %v", result)
	}

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}

func TestProjectorReturnsTransformedElements(t *testing.T) {
	many := &many{items: reflect.ValueOf([]string{"a", "b", "c"})}
	projector := projector{
		source:      many.Source,
		projections: []reflect.Value{reflect.ValueOf(func(x string) string { return strings.ToUpper(x) })},
	}

	var s WrappedSource = projector.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "A" {
		t.Fatalf("Expected element 'A' but got %v", result)
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "B" {
		t.Fatalf("Expected element 'B' but got %v", result)
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "C" {
		t.Fatalf("Expected element 'C' but got %v", result)
	}

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}

func TestProjectorSupportsMultipleReturnValues(t *testing.T) {
	many := &many{items: reflect.ValueOf([]string{"a", "b", "c"})}
	projector := &projector{
		source:      many.Source,
		projections: []reflect.Value{reflect.ValueOf(func(x string) (string, string) { return strings.ToUpper(x), x + "X" })},
	}

	var s WrappedSource = projector.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "A" {
		t.Fatalf("Expected element 'A' but got %v", result)
	}
	if result[1].String() != "aX" {
		t.Fatalf("Expected element 'aX' but got %v", result)
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a second element but none come next")
	}
	if result[0].String() != "B" {
		t.Fatalf("Expected element 'B' but got %v", result)
	}
	if result[1].String() != "bX" {
		t.Fatalf("Expected element 'bX' but got %v", result)
	}

	result, s = s()

	if s == nil {
		t.Fatal("Expected a third element but none come next")
	}
	if result[0].String() != "C" {
		t.Fatalf("Expected element 'C' but got %v", result)
	}
	if result[1].String() != "cX" {
		t.Fatalf("Expected element 'cX' but got %v", result)
	}

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
