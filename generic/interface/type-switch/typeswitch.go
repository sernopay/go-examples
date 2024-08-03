package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("the type is %T value %d \n", v, v)
	case string:
		fmt.Printf("the type is %T value %s \n", v, v)
	default:
		fmt.Println("I Don't Know The Type")
	}
}

func main() {
	do("Hello")
	do(35)
	do(true)
}
