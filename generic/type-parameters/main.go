package main

import "fmt"

func Index[T comparable](list []T, val T) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}
	return -1
}

func main() {
	// Index func works on slice of int
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Index(ints, 3))

	// Index func also works on slice of string
	ss := []string{"a", "b", "c", "d", "e"}
	fmt.Println(Index(ss, "a"))
}
