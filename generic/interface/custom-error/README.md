# Custom Error Example

This example demonstrates how to define and use custom error types in Go. The code defines a `CustomError` type that implements the `error` interface and uses it in a simple validation function.

## Code Explanation

### Defining the CustomError Type
The `CustomError` type is defined as an empty struct. It implements the `error` interface by providing an `Error` method that returns a custom error message.
```go
type CustomError struct {
}

func (e *CustomError) Error() string {
    return "This is a custom error"
}
```

### Implementing the validateInput Function
The `validateInput` function takes an `interface{}` as input and returns an `error`. If the input is `nil`, it returns an instance of `CustomError`. Otherwise, it returns `nil`.
```go
func validateInput(v interface{}) error {
    if v == nil {
        return &CustomError{}
    }
    return nil
}
```

### Main Function
In the `main` function, the `validateInput` function is called with a `nil` argument. If an error is returned and it is of type `CustomError`, it is printed to the console.
```go
func main() {
    err := validateInput(nil)
    if err, ok := err.(*CustomError); ok {
        fmt.Println(err)
    }
}
```

## Running the Code
To run this code, execute the following command in your terminal:
```bash
go run error.go
```

You should see the following output:
```
This is a custom error
```

## Summary
This example shows how to define a custom error type in Go, implement the error interface, and use the custom error in a function. Custom error types can provide more context and information about errors, making it easier to debug and handle errors in your programs.