# Type Parameters Example
This example demonstrates how to use type parameters in Go. In this example, we will create an `Index` function that returns the index of an element in a slice of any type. If the element is not found, the function will return -1.

## About Type Parameters
Type parameters allow you to define functions, methods, that can work with any type. Or we can call it generic functions. This is useful when you want to write generic code that can be used with different types without having to duplicate the code for each type. The type parameters of a function appear between brackets before the function's arguments. For example
```go
func Index[T comparable](list []T, val T) int
```
This declaration means that `list` is a slice of any type `T` that fulfills the built-in constraint comparable. `val` is also a value of the same type.

## Code Explanation

### Generic Index Function:
```go
func Index[T comparable](list []T, val T) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}
	return -1
}
```
* The Index function is defined with a type parameter `T` that is constrained by the `comparable` interface. This means that `T` can be any type that supports comparison operations (e.g., `==` and `!=`).
* The function takes two parameters: a slice of type `T` (`list`) and a value of type `T` (`val`).
* It iterates over the slice using a for loop. For each element, it checks if the element is equal to the value (`val`).
* If a match is found, it returns the index of the element.
* If no match is found after iterating through the entire slice, it returns `-1`.

### Main Function:
The main function demonstrates how to use the Index function with slices of different types.
```go
func main() {
	// Index func works on slice of int
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Index(ints, 3))

	// Index func also works on slice of string
	ss := []string{"a", "b", "c", "d", "e"}
	fmt.Println(Index(ss, "a"))
}
```
#### Using Index with a Slice of Integers:
* We create a slice of integers `ints` and call the `Index` function to find the index of the value `3`.
* The function returns `3`, which is the index of the value `3` in the slice.

#### Using Index with a Slice of Strings:
* We create a slice of strings `ss` and call the `Index` function to find the index of the value `"a"`.
* The function returns `0`, which is the index of the value `"a"` in the slice.

## Run the Example
To run the example, use the following command:
```bash
go run main.go
```
The output will be:
```bash
3
0
```

## Summary
This example demonstrates how to use type parameters in Go to create generic functions that can work with any type. This is useful when you want to write generic code that can be used with different types without having to duplicate the code for each type.