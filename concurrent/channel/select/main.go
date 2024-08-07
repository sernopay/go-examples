package main

import "fmt"

func group1(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
}

func group2(c chan int) {
	for i := 6; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	cg1 := make(chan int)
	cg2 := make(chan int)

	go group1(cg1)
	go group2(cg2)

	for {
		select {

		case val, ok := <-cg1:
			if ok {
				fmt.Println("print from cg1 ", val)
			} else {
				cg1 = nil
			}

		case val, ok := <-cg2:
			if ok {
				fmt.Println("print from cg2 ", val)
			} else {
				cg2 = nil
			}
		}

		if cg1 == nil && cg2 == nil {
			break
		}
	}
}
