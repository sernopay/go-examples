# Custom Error Example

This example demonstrates how to create and use custom error types in Go. We will explore two examples of custom error handling: one using a custom error type and another using the `errors.New` function.

## Example 1: Custom Error Type
In the first example, we define a custom error type by implementing the error interface. This involves creating a struct and defining an `Error` method for it.

```go
type errorCustom struct {
}

func (e *errorCustom) Error() string {
    return "this is a custom error"
}

var EustomErrInstance = &errorCustom{}

func validateInput(v interface{}) error {
    if v == nil {
        return EustomErrInstance
    }
    return nil
}
```
### Explanation
1. **Custom Error Type**: We define a struct errorCustom and implement the Error method to return a custom error message.
2. **Error Instance**: We create an instance of the custom error type. This instance is public and can be used in other packages, so it will make it easier to compare errors later.
3. **Validation Function**: The validateInput function returns the custom error instance if the input is nil.

## Example 2: Using errors.New
In the second example, we use the `errors.New` function to create a custom error message.
```go
var ErrorCustom2 = errors.New("this is a custom error 2")

func validateInput2(v interface{}) error {
    if v == nil {
        return ErrorCustom2
    }
    return nil
}
```
### Explanation
1. **Custom Error Message**: We use `errors.New` to create a custom error message. Same as the previous example, this error is public and can be used in other packages. So, it will make it easier to compare errors later.
2. **Validation Function**: The validateInput2 function returns the custom error if the input is nil.

## Main Function
The main function demonstrates how to use the custom error handling mechanisms defined in the previous examples.
```go
func main() {
    err := validateInput(nil)
    if err == EustomErrInstance {
        fmt.Println(err)
    }

    err2 := validateInput2(nil)
    if err2 == ErrorCustom2 {
        fmt.Println(err2)
    }
}
```
Note that by using public error instances, we can compare errors directly without having to compare error messages or use type assertions.

## Running the Code
To run this code, execute the following command in your terminal:
```bash
go run error.go
```

You should see the following output:
```
this is a custom error
this is a custom error 2
```

## When to Implement Custom Error Types or Use errors.New
- **Custom Error Types**: Use custom error types when you need to provide additional context or information about the error, for example, error codes, error details, or error types. This can be useful for debugging and handling errors in a more structured way.
- **errors.New**: Use `errors.New` when you need a simple custom error message without additional context. This is useful for creating custom error messages quickly without defining a new error type.

## Summary
This example shows how to define a custom error type in Go, implement the error interface, and use the custom error in a function. Custom error types can provide more context and information about errors, making it easier to debug and handle errors in your programs.