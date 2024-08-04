package main

import "fmt"

type CustomError struct {
}

func (e *CustomError) Error() string {
	return "This is a custom error"
}

func validateInput(v interface{}) error {
	if v == nil {
		return &CustomError{}
	}
	return nil
}

func main() {
	err := validateInput(nil)
	if err, ok := err.(*CustomError); ok {
		fmt.Println(err)
	}
}
