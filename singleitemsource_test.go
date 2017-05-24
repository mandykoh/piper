package piper

import "testing"

func TestSingleItemSourceReturnsValuesForOneItem(t *testing.T) {
	s := singleItemSource{values: []interface{}{"a", 1}}

	result, ok := s.Next()

	if !ok {
		t.Fatal("Expected an element but none come next")
	}
	if result[0].String() != "a" {
		t.Fatalf("Expected element 'a' but got %v", result[0])
	}
	if result[1].Int() != 1 {
		t.Fatalf("Expected element 1 but got %v", result[1])
	}

	result, ok = s.Next()

	if ok {
		t.Fatalf("Expected no more elements but got %v", result)
	}
}
