package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

// Push adds an element to the stack.
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var elementZeroValues T
		return elementZeroValues, false
	}
	return s.elements[len(s.elements)-1], true
}

func main() {
	// Create a stack of integers
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(3)
	intStack.Push(5)

	fmt.Println(intStack.Peek())

	// Create a stack of strings
	stringStack := &Stack[string]{}
	stringStack.Push("a")
	stringStack.Push("c")
	stringStack.Push("e")
	fmt.Println(stringStack.Peek())
}
