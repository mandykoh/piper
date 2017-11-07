package piper

import (
	"reflect"
	"testing"
)

func TestAggregationConsumesSourceElements(t *testing.T) {
	many := &many{items: reflect.ValueOf([]string{"a", "b", "c"})}
	a := aggregation{
		source:     many.Source,
		seed:       reflect.ValueOf("z"),
		aggregator: reflect.ValueOf(func(result, x string) string { return result + x }),
	}

	var s WrappedSource = a.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if count := len(result); count != 1 {
		t.Fatalf("Expected one value but got %d", count)
	}
	if result[0].String() != "zabc" {
		t.Errorf("Expected element 'zabc' but got %v", result)
	}

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}

func TestAggregationUsesRuntimeTypeOfElement(t *testing.T) {
	many := &many{items: reflect.ValueOf([]interface{}{"a", "b", "c"})}
	a := aggregation{
		source:     many.Source,
		seed:       reflect.ValueOf(""),
		aggregator: reflect.ValueOf(func(result, x string) string { return result + x }),
	}

	var s WrappedSource = a.Source
	var result []reflect.Value

	result, s = s()

	if s == nil {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "abc" {
		t.Fatalf("Expected element 'abc' but got %v", result[0])
	}

	result, s = s()

	if s != nil {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
