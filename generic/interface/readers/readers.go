package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	data := "Hello World"
	reader := strings.NewReader(data)

	buf := make([]byte, 4)

	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n is %d | ", n)
		fmt.Printf("the buf now contains %s \n", string(buf[:n]))
	}
}
