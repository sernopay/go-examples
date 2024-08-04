# Reader Example

This example demonstrates how to use the `Reader` interface to read data. In this example, we will read data in chunks from a string using the `strings.NewReader` function and the `Read` method.

## Code Explanation
The following Go code reads a string in chunks of 4 bytes and prints each chunk:
```go
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
```

### Detailed Explanation
#### Data Initialization
We initialize a string `data` with the value `"Hello World"`.
```go
data := "Hello World"
```

#### Creating a Reader
We create a new reader from the string using `strings.NewReader(data)`. This reader implements the `io.Reader interface`.
```go
reader := strings.NewReader(data)
```

#### Buffer Initialization
We create a buffer `buf` of size 4 bytes. This buffer will hold the chunks of data read from the reader.
```go
buf := make([]byte, 4)
```

#### Reading Data in a Loop
```go
for {
    n, err := reader.Read(buf)
    if err == io.EOF {
        break
    }
    fmt.Printf("n is %d | ", n)
    fmt.Printf("the buf now contains %s \n", string(buf[:n]))
}
```
* We use a for loop to read data from the reader in chunks.
* Inside the loop, we call `reader.Read(buf)` to read up to 4 bytes of data into the buffer.
* The `Read` method returns the number of bytes read (`n`) and an error (`err`).
* If the error is `io.EOF`, it means we have reached the end of the data, and we break out of the loop.
* We print the number of bytes read and the contents of the buffer up to `n` bytes.

## Running the Code
To run the code execute it using the `go run` command:
```bash
go run .
```

the output will be:
```bash
n is 4 | the buf now contains Hell 
n is 4 | the buf now contains o Wo 
n is 3 | the buf now contains rld
```

## Summary
This example demonstrates how to use the io.Reader interface to read data in fixed-size chunks. This approach is **useful when dealing with large data streams where reading the entire data at once is not feasible**. By reading data in chunks, you can process the data incrementally, which is more memory-efficient and can handle larger data sizes.