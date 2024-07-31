# Declare a type constraint

## Overview

This Go program demonstrates the use of type constraints with generics to sum values in a map. The program defines a generic function `sumNumber` that can sum the values of a map with keys of any comparable type and values of either `int64` or `float64`.

## Type Constraint
```
type Number interface {
    int64 | float64
}
```
A type constraint `Number` is defined using a type set. It specifies that the `Number` interface can be either `int64` or `float64`.

## Generic Function
```
func sumNumber[K comparable, V Number](m map[K]V) V {
    var total V
    for _, v := range m {
        total += v
    }
    return total
}
```
The `sumNumber` function is a generic function that takes a map `m` with keys of any comparable type `K` and values of type `V` (which must satisfy the `Number` constraint). It iterates over the map, sums the values, and returns the total.

## Main Function
```
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
```

The main function demonstrates the usage of the `sumNumber` function with two maps:
* ints: A map with string keys and int64 values.
* floats: A map with string keys and float64 values.

## Output
When you run this program, it will output:
```
Print sum ints 3 
Print sum floats 7.700000
```

## Summary
This demonstrates the flexibility and power of Go's generics and type constraints.