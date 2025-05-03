# Generic Types Example

## Learning Objectives
* Understand the concept of generic types in Go.
* Understand the benefits of using generics in Go.
* Learn how to define and use generic types in Go.

## About Generic Types
Generics were introduced in Go 1.18, allowing developers to write code that can work with any data type. This means you can create functions and data structures that are not limited to a specific type, making your code more flexible and reusable. Generic types are defined using type parameters, which are placeholders for the actual types that will be used when the code is executed.

### Why Use Generics?
* Generics allow you to write code that is more flexible and reusable. Instead of writing separate implementations for each type, you can write a single implementation that works with any type. Therefore, you can avoid code duplication and make your code more maintainable.
* Generics also provide type safety. When you use a generic type, the compiler checks that the types you use are compatible with the operations you perform on them. This helps to catch errors at compile time rather than runtime.

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
		return *new(T), false
	}
	return s.elements[len(s.elements)-1], true
}

...

// IsEmpty checks whether the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
```
* We define a generic `Stack` type using a type parameter `T`. T can be any type, including built-in types like `int`, `string`, or user-defined types. The `Stack` struct contains a slice of elements of type `T`. 
* The `Push` method adds an element to the stack. It appends the element to the `elements` slice.
* The `Peek` method returns the top element of the stack without removing it. It returns the top element and a boolean indicating whether the operation was successful. If the stack is empty, it returns a zero value of type `T` and `false`.

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
* From this example, we can see that the `Stack` type is generic and can be used with different types. We created two stacks: one for integers and one for strings. The `Push` and `Peek` methods work with both stacks, demonstrating the flexibility of generics. By doing this, we avoid code duplication and make our code more maintainable.
* We can't use the same `Stack` type for both integers and strings, demonstrating the type safety of generics.

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

For full code example, you can find it in this here