# piper

[![GoDoc](https://godoc.org/github.com/mandykoh/piper?status.svg)](https://godoc.org/github.com/mandykoh/piper)
[![Go Report Card](https://goreportcard.com/badge/github.com/mandykoh/piper)](https://goreportcard.com/report/github.com/mandykoh/piper)
[![Build Status](https://travis-ci.org/mandykoh/piper.svg?branch=master)](https://travis-ci.org/mandykoh/piper)

Go library for lazily-evaluated pipeline processing.


## Introduction

Piper uses constructs inspired by [LINQ](https://en.wikipedia.org/wiki/Language_Integrated_Query) to make it easy to build lazily-evaluated, pipeline style processing code.

As Go doesn’t have generics, Piper is reflection based, and thus loses some static type checking/inference to do what it does.


## Examples

### Sources

Trivially pipe a single value to `Println()`:

```go
piper.FromSingle("Hello, World!").To(fmt.Println)

// Outputs:
// Hello, World!
```

Pipe a single value to `Printf()`:

```go
piper.FromSingle("Hello, World!").To(func(s string) { fmt.Printf("%s\n", s) })

// Outputs:
// Hello, World!
```

Pipe a single pair of values to `Printf()` (note that the multiple values becomes the input to the next pipeline stage):

```go
piper.FromSingle("Hello", "World").To(func(greeting, subject string) { fmt.Printf("%s, %s!\n", greeting, subject) })

// Outputs:
// Hello, World!
```

Stream a few values from a slice to `Println()`:

```go
piper.FromMany([]string{ "apple", "pear", "banana" }).To(fmt.Println)

// Outputs:
// apple
// pear
// banana
```

Stream values from a custom source function:

```go
type CountDownSource func() (value int, restOrEnd CountDownSource)

func countDown(n int) CountDownSource {
  return func() (int, CountDownSource) {
    if n < 0 {
      return 0, nil
    }
    return n, countDown(n - 1)
  }
}

piper.From(countDown(3)).To(fmt.Println)

// Outputs:
// 3
// 2
// 1
// 0
```

The source function must return a value and a function which produces the following value. If there is no value to return, `nil` should be returned in place of the function.


### Filtering

Exclude words containing the letter `e` using a `Where` filter:

```go
piper.FromMany([]string{ "apple", "pear", "banana" }).
  Where(func(s string) bool { return !strings.Contains(s, "e") }).
  To(fmt.Println)

// Outputs:
// banana
```

Make the above more readable using a helper partial function:

```go
func StringExcludes(s string) func(string) bool {
  return func(v string) bool {
    return !strings.Contains(v, s)
  }
}

piper.FromMany([]string{ "apple", "pear", "banana" }).
  Where(StringExcludes("e")).
  To(fmt.Println)
```

### Projection

Get each word and its length using a `Select` projection (note that the return type of the function passed to `Select` becomes the input to the next stage of the pipe):

```go
piper.FromMany([]string{ "apple", "pear", "banana" }).
  Select(func(s string) (string, int) { return s, len(s) }).
  To(func(s string, l int) { fmt.Printf("String: %s, Length: %d\n", s, l) })

// Outputs:
// String: apple, Length: 5
// String: pear, Length: 4
// String: banana, Length: 6
```

Decode a base64-encoded string using a `Select` projection with multiple return values:

```go
piper.FromSingle("SGVsbG8sIFdvcmxkIQ==").
  Select(base64.StdEncoding.DecodeString).
  Where(func(decoded []byte, err error) bool { return err == nil }).
  Select(func(decoded []byte, err error) string { return string(decoded) }).
  To(fmt.Println)

// Outputs:
// Hello, World!
```

Get the upper and lower case version of each word using multiple projections with one `Select`:

```go
piper.FromMany([]string{"apple", "pear", "banana"}).
  Select(strings.ToUpper, strings.ToLower).
  To(func(s1, s2 string) { fmt.Printf("%s %s\n", s1, s2) })

// Outputs:
// APPLE apple
// PEAR pear
// BANANA banana
```

### Aggregation

Count items using an `Aggregate`:

```go
piper.FromMany([]string{ "apple", "pear", "banana" }).
  Aggregate(0, func(total, item string) int { return total + 1 }).
  To(fmt.Println)

// Outputs:
// 3
```

### Flattening

Turn a single slice or array value (eg a slice of words) into multiple values (eg a one-word-at-a-time stream) using a `Flatten` pipeline stage:

```go
piper.FromSingle([]string{ "apple", "pear", "banana" }).
  Flatten().
  To(fmt.Println)

// Outputs:
// apple
// pear
// banana
```

Combine a value with each word in a slice by `Flatten`ing:

```go
piper.FromSingle("Hello", []string{ "apple", "pear", "banana" }).
  Flatten().
  To(func(greeting, fruit string) { fmt.Printf("%s, %s\n!", greeting, fruit) })

// Outputs:
// Hello, apple!
// Hello, pear!
// Hello, banana!
```

Flattening multiple slices/arrays produces a Cartesian join. To find all combinations of sizes, colours, and shapes:

```go
sizes := []string{"small", "large"}
colours := []string{"blue", "red", "green"}
shapes := []string{"square", "circle", "triangle"}

piper.FromSingle(sizes, colours, shapes).
  Flatten().
  To(func(size, colour, shape string) { fmt.Printf("%s %s %s\n", size, color, shape) })

// Outputs:
// small blue square
// small blue circle
// small blue triangle
// small red square
// small red circle
// ...etc
```

### More examples

Combine multiple pipeline stages for more complex processing:

```go
piper.FromSingle([]Person{
    {FirstName: "John", LastName: "Doe"},
    {FirstName: "Jane", LastName: "Dee"},
    {FirstName: "Bob", LastName: "Smith"},
  }).
  Flatten().
  Select(func(p Person) string { return p.FirstName + " " + p.LastName }).
  Where(func(fullName string) bool { return len(fullName) < 9 }).
  To(fmt.Println)

// Outputs the full names that are less than 9 characters long:
// John Doe
// Jane Dee
```

Make the above more readable by defining helper partial functions:

```go
func FullName(p Person) string {
  return p.FirstName + " " + p.LastName
}

func LengthLessThan(max int) func(s string) bool {
  return func(s string) bool {
    return len(s) < max
  }
}

piper.FromSingle([]Person{
    {FirstName: "John", LastName: "Doe"},
    {FirstName: "Jane", LastName: "Dee"},
    {FirstName: "Bob", LastName: "Smith"},
  }).
  Flatten().
  Select(FullName).
  Where(LengthLessThan(9)).
  To(fmt.Println)
```

Parse some JSON using a helper function:

```go
func jsonProperty(name string) func(map[string]interface{}) interface{} {
  return func(jsonObject map[string]interface{}) interface{} {
    return jsonObject[name]
  }
}

piper.FromSingle(jsonData).
  Select(jsonProperty("items")).
  Flatten().
  Select(jsonProperty("id"), jsonProperty("url")).
  To(func(id, url string) {
    fmt.Printf("%s: %s\n", id, url)
  })
```
