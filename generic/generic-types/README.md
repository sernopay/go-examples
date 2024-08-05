# Generic Types Example

This example demonstrates how to use generic types.

## About Generic Types
A type can be parameterized with a type parameter, which could be useful for implementing generic data structures. For example, a generic stack.

## Code Explanation

### Generic Stack
The following Go code defines a generic `Stack` type that can store elements of any type.
```go
type Stack[T any] struct {
    elements []T
}

// Push adds an element to the stack.
func (s *Stack[T]) Push(element T) {
    s.elements = append(s.elements, element)
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
    if len(s.elements) == 0 {
        var elementZeroValues T
        return elementZeroValues, false
    }
    return s.elements[len(s.elements)-1], true
}
```
* We define a generic `Stack` type using a type parameter `T`. The `Stack` struct contains a slice of elements of type `T`.
* The `Push` method adds an element to the stack. It appends the element to the `elements` slice.
* The `Peek` method returns the top element of the stack without removing it. It returns the top element and a boolean value indicating whether the stack is empty.

### main function
The `main` function demonstrates the usage of the generic `Stack` type.
```go
func main() {
    // Create a stack of integers
    intStack := &Stack[int]{}
    intStack.Push(1)
    intStack.Push(3)
    intStack.Push(5)

    fmt.Println(intStack.Peek()) // Output: 5 true

    // Create a stack of strings
    stringStack := &Stack[string]{}
    stringStack.Push("a")
    stringStack.Push("c")
    stringStack.Push("e")
    fmt.Println(stringStack.Peek()) // Output: e true
}
```
* We create a stack of integers and push elements `1`, `3`, and `5` to the stack. We then print the top element of the stack using the `Peek` method.
* We create a stack of strings and push elements `"a"`, `"c"`, and `"e"` to the stack. We then print the top element of the stack using the `Peek` method.

## Run the code
To run the code, use the following command:
```bash
go run main.go
```
The output will be:
```bash
5 true
e true
```

## Summary
This example demonstrates how to use generic types in Go. We defined a generic `Stack` type that can store elements of any type and demonstrated its usage with integers and strings, making it versatile and reusable data structure. By using generics, we can write more flexible and type-safe code that can handle different data types without duplication.
