# piper
Go library for lazily-evaluated pipeline processing.

## Introduction

Piper uses constructs inspired by [LINQ](https://en.wikipedia.org/wiki/Language_Integrated_Query) to make it easy to build lazily-evaluated, pipeline style processing code.

As Go doesn’t have generics, Piper is reflection based, and thus loses some static type checking/inference to do what it does.


## Examples

Trivially pipe a single value to `Println()`:

```go
piper.From("Hello, World!").To(fmt.Println)

// Outputs:
// Hello, World!
```

Pipe a single value to `Printf()`:

```go
piper.From("Hello, World!").To(func(s string) { fmt.Printf("%s\n", s) })

// Outputs:
// Hello, World!
```

Pipe a single pair of values to `Printf()` (note that the multiple values becomes the input to the next pipeline stage):

```go
piper.From("Hello", "World").To(func(greeting, subject string) { fmt.Printf("%s, %s!\n", greeting, subject) })

// Outputs:
// Hello, World!
```

Stream a few values from a slice to `Println()`:

```go
piper.FromIndexable([]string{ "apple", "pear", "banana" }).To(fmt.Println)

// Outputs:
// apple
// pear
// banana
```

Exclude words containing the letter `e` using a `Where` filter:

```go
piper.FromIndexable([]string{ "apple", "pear", "banana" }).
  Where(func(s string) bool { return !strings.Contains(s, "e") }).
  To(fmt.Println)

// Outputs:
// banana
```

Get each word and its length using a `Select` projection (note that the return type of the function passed to `Select` becomes the input to the next stage of the pipe):

```go
piper.FromIndexable([]string{ "apple", "pear", "banana" }).
  Select(func(s string) (string, int) { return s, len(s) }).
  To(func(s string, l int) { fmt.Printf("String: %s, Length: %d\n", s, l) })

// Outputs:
// String: apple, Length: 5
// String: pear, Length: 4
// String: banana, Length: 6
```

Decode a base64-encoded string using a `Select` projection with multiple return values:

```go
piper.From("SGVsbG8sIFdvcmxkIQ==").
  Select(base64.StdEncoding.DecodeString).
  Where(func(decoded []byte, err error) bool { return err == nil }).
  Select(func(decoded []byte, err error) string { return string(decoded) }).
  To(fmt.Println)

// Outputs:
// Hello, World!
```

Get the upper and lower case version of each word using multiple projections with one `Select`:

```go
piper.FromIndexable([]string{"apple", "pear", "banana"}).
  Select(strings.ToUpper, strings.ToLower).
  To(func(s1, s2 string) { fmt.Printf("%s %s\n", s1, s2) })

// Outputs:
// APPLE apple
// PEAR pear
// BANANA banana
```

Turn a single slice or array value (eg a slice of words) into multiple values (eg a one-word-at-a-time stream) using a `Flatten` pipeline stage:

```go
piper.From([]string{ "apple", "pear", "banana" }).
  Flatten().
  To(fmt.Println)

// Outputs:
// apple
// pear
// banana
```

Combine a value with each word in a slice by `Flatten`ing:

```go
piper.From("Hello", []string{ "apple", "pear", "banana" }).
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

piper.From(sizes, colours, shapes).
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

Combine multiple pipeline stages for more complex processing:

```go
piper.From([]Person{
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

piper.From([]Person{
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

piper.From(jsonData).
  Select(jsonProperty("items")).
  Flatten().
  Select(jsonProperty("id"), jsonProperty("url")).
  To(func(id, url string) {
    fmt.Printf("%s: %s\n", id, url)
  })
```
