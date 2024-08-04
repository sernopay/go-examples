package main

import (
	"errors"
	"fmt"
)

// example 1
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

// example 2
var ErrorCustom2 = errors.New("this is a custom error 2")

func validateInput2(v interface{}) error {
	if v == nil {
		return ErrorCustom2
	}
	return nil
}

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
