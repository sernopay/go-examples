# Type Assertion Example

This Go program demonstrates the concept of type assertion, which is used to extract the underlying value of an interface variable. Below is a detailed explanation of the code.

## Code Explanation

### Declaring an Interface Variable:
```go
var i interface{} = "hello"
```
Here, `i` is declared as an empty interface (`interface{}`), which means it can hold a value of any type. In this case, it is assigned the string `"hello"`.

### Type Assertion:
```go
if s, ok := i.(string); ok {
    fmt.Println(s)
}
```
* The type assertion `i.(string)` checks if the interface value `i` holds a string.
* If the assertion is successful, `s` will hold the string value, and `ok` will be `true`.
* If the assertion fails, `ok` will be `false`, and the code inside the if block will not execute.

## Running the Program
To run the program, use the following command:
```bash
go run assertion.go
```
This will output:
```
hello
```