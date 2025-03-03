# Variadic Parameters Example
This example demonstrates how to use variadic parameters in Go functions. 

## About Variadic Parameters
Variadic parameters allows you to pass an arbitrary number of values of type T to a function. This is useful when the number of arguments is not known beforehand.

## Why use variadic parameters?
- Simplifies handling multiple parameters without needing to define slices explicitly.
- More flexible API design (e.g., logging functions that accept multiple values).
- Useful for wrapper functions where arguments are forwarded to another function.

## Code Explanation
### Variadic Sum Function:
```go
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
```
This declaration means that the `sum` function can take any number of integer arguments. The `numbers` parameter is a slice of integers that can be accessed like a regular slice.

### The main Function
The main function demonstrates how to use the `sum` function with different numbers of arguments.
```go
func main() {
	fmt.Println(sum(1, 2, 3)) // 6
	fmt.Println(sum(10, 20))  // 30
}
```
#### Using sum with Multiple Arguments:
* We call the `sum` function with multiple arguments to calculate the sum of those numbers.
* The sum parameter can accept any number of integer arguments. For example in the first call, we pass 3 arguments (1, 2, 3) and in the second call, we pass 2 arguments (10, 20).

## Run the Example
To run the example, use the following command:
```bash
go run main.go
```
The output will be:
```bash
6
30
```

## Summary
Variadic parameters in Go allow you to pass an arbitrary number of values of type T to a function. This is useful when the number of arguments is not known beforehand. In this example we created a `sum` function that calculates the sum of multiple integers passed as arguments. The function can accept any number of integer arguments and returns the total sum of those numbers. 