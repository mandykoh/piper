package piper

import "testing"

func TestPipeDraining(t *testing.T) {
	var results []string

	FromIndexable([]string{"a", "b", "c"}).
		To(func(v string) { results = append(results, v) })

	if count := len(results); count != 3 {
		t.Fatalf("Expected 3 elements but got %d", count)
	}
	if results[0] != "a" {
		t.Errorf("Expected element 'a' but got %v", results[0])
	}
	if results[1] != "b" {
		t.Errorf("Expected element 'b' but got %v", results[1])
	}
	if results[2] != "c" {
		t.Errorf("Expected element 'c' but got %v", results[2])
	}
}

func TestPipeDrainingWithFiltering(t *testing.T) {
	var results []string

	FromIndexable([]string{"a", "b", "c"}).
		Where(func(v string) bool { return v != "b" }).
		To(func(v string) { results = append(results, v) })

	if count := len(results); count != 2 {
		t.Fatalf("Expected 2 elements but got %d", count)
	}
	if results[0] != "a" {
		t.Errorf("Expected element 'a' but got %v", results[0])
	}
	if results[1] != "c" {
		t.Errorf("Expected element 'c' but got %v", results[1])
	}
}

func TestPipeDrainingWithProjection(t *testing.T) {
	var results1 []int
	var results2 []int

	FromIndexable([]int{1, 2, 3}).
		Select(func(n int) (int, int) { return n, n * 2 }).
		To(func(n1, n2 int) { results1 = append(results1, n1); results2 = append(results2, n2) })

	if count := len(results1); count != 3 {
		t.Fatalf("Expected 3 elements but got %d", count)
	}
	if results1[0] != 1 {
		t.Errorf("Expected element 1 but got %v", results1[0])
	}
	if results1[1] != 2 {
		t.Errorf("Expected element 2 but got %v", results1[1])
	}
	if results1[2] != 3 {
		t.Errorf("Expected element 3 but got %v", results1[2])
	}

	if count := len(results2); count != 3 {
		t.Fatalf("Expected 3 elements but got %d", count)
	}
	if results2[0] != 2 {
		t.Errorf("Expected element 2 but got %v", results2[0])
	}
	if results2[1] != 4 {
		t.Errorf("Expected element 4 but got %v", results2[1])
	}
	if results2[2] != 6 {
		t.Errorf("Expected element 6 but got %v", results2[2])
	}
}
