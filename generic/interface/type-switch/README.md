# Type Switch Example

This Go program demonstrates the use of a type switch to handle different types stored in an interface variable. Below is a detailed explanation of the code.

## Code Explanation

### do Function

#### Function Declaration:
```go
func do(i interface{}) {
```
* The do function takes a single parameter i of type interface{}, which means it can accept any type of value.

#### Type Switch:
```go
switch v := i.(type) {
```
* A type switch is used to determine the type of the value stored in the interface `i`.
* The syntax `v := i.(type)` assigns the value of `i` to `v` and checks its type.

#### Case for int:
```go
case int:
    fmt.Printf("the type is %T value %d \n", v, v)
```
* If `i` is of type `int`, it prints the type and value using `fmt.Printf`.

#### Case for string:
```go
case string:
    fmt.Printf("the type is %T value %s \n", v, v)
```
* If `i` is of type `string`, it prints the type and value using `fmt.Printf`.

#### Default Case:
```go
default:
    fmt.Println("I Don't Know The Type")
```
* If `i` is of any other type, it prints a message indicating that the type is unknown.

## Running the Program
To run the program, use the following command:
```bash
go run typeswitch.go
```

This will output:
```
the type is string value Hello 
the type is int value 35 
I Don't Know The Type
```