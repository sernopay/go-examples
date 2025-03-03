package main

import "fmt"

// variadic parameters that accepts any number of int arguments
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3)) // 6
	fmt.Println(sum(10, 20))  // 30
}
