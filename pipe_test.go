package piper

import "testing"

func TestPipeAggregation(t *testing.T) {
	var results []int

	items := []int{1, 2, 3, 4}

	FromMany(items).
		Aggregate(0, func(total, n int) int { return total + n }).
		To(func(n int) { results = append(results, n) })

	if count := len(results); count != 1 {
		t.Fatalf("Expected 1 element but got %d", count)
	}
	if results[0] != 10 {
		t.Errorf("Expected value 10 but got %v", results[0])
	}
}

func TestPipeCreationUsingCustomFunction(t *testing.T) {
	var results []int

	From(countDown(3)).To(func(n int) { results = append(results, n) })

	if count := len(results); count != 4 {
		t.Fatalf("Expected 4 elements but got %d", count)
	}
	if results[0] != 3 {
		t.Errorf("Expected element 3 but got %v", results[0])
	}
	if results[1] != 2 {
		t.Errorf("Expected element 2 but got %v", results[1])
	}
	if results[2] != 1 {
		t.Errorf("Expected element 1 but got %v", results[2])
	}
	if results[3] != 0 {
		t.Errorf("Expected element 0 but got %v", results[3])
	}
}

func TestPipeCreationUsingCustomFunctionWithNoResults(t *testing.T) {
	var results []int

	From(countDown(-1)).To(func(n int) { results = append(results, n) })

	if count := len(results); count != 0 {
		t.Fatalf("Expected no elements but got %d", count)
	}
}

func TestPipeDraining(t *testing.T) {
	var results []string

	FromMany([]string{"a", "b", "c"}).
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

	FromMany([]string{"a", "b", "c"}).
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

func TestPipeDrainingWithFlattening(t *testing.T) {
	var results []string

	FromSingle([]string{"blue", "red", "green"}, []string{"square", "circle", "triangle"}).
		Flatten().
		To(func(color, shape string) { results = append(results, color+" "+shape) })

	if count := len(results); count != 9 {
		t.Fatalf("Expected 9 elements but got %d", count)
	}
	if expected := "blue square"; results[0] != expected {
		t.Errorf("Expected %v but got %v", expected, results[0])
	}
	if expected := "blue circle"; results[1] != expected {
		t.Errorf("Expected %v but got %v", expected, results[1])
	}
	if expected := "blue triangle"; results[2] != expected {
		t.Errorf("Expected %v but got %v", expected, results[2])
	}
	if expected := "red square"; results[3] != expected {
		t.Errorf("Expected %v but got %v", expected, results[3])
	}
	if expected := "red circle"; results[4] != expected {
		t.Errorf("Expected %v but got %v", expected, results[4])
	}
	if expected := "red triangle"; results[5] != expected {
		t.Errorf("Expected %v but got %v", expected, results[5])
	}
	if expected := "green square"; results[6] != expected {
		t.Errorf("Expected %v but got %v", expected, results[6])
	}
	if expected := "green circle"; results[7] != expected {
		t.Errorf("Expected %v but got %v", expected, results[7])
	}
	if expected := "green triangle"; results[8] != expected {
		t.Errorf("Expected %v but got %v", expected, results[8])
	}
}

func TestPipeDrainingWithProjection(t *testing.T) {
	var results1 []int
	var results2 []int

	FromMany([]int{1, 2, 3}).
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

type CountDownSource func() (value int, restOrEnd CountDownSource)

func countDown(n int) CountDownSource {
	return func() (int, CountDownSource) {
		if n < 0 {
			return 0, nil
		}
		return n, countDown(n - 1)
	}
}
