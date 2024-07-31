package main

import "fmt"

type Number interface {
	int64 | float64
}

func sumNumber[K comparable, V Number](m map[K]V) V {
	var total V
	for _, v := range m {
		total += v
	}
	return total
}

func main() {
	ints := map[string]int64{
		"first":  1,
		"second": 2,
	}

	floats := map[string]float64{
		"first":  3.3,
		"second": 4.4,
	}

	fmt.Printf("Print sum ints %d \n", sumNumber(ints))
	fmt.Printf("Print sum floats %f \n", sumNumber(floats))
}
