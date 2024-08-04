package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

func (e *Employee) String() string {
	return fmt.Sprintf("Employee name: %s age: %d", e.Name, e.Age)
}

func main() {
	e := &Employee{"John", 25}
	fmt.Println(e)
}
