# Create Module

go mod enables dependency tracking for your code.

## Example of how to create a module
If you publish a module, this must be a path from which your module can be downloaded by Go tools. That would be your code's repository.
```
go mod init example.com/greetings
```
The go mod init command creates a go.mod file to track your code's dependencies.

The previous example has the following structure: :
```
<home>/
 |-- greetings/
    |-- greetings.go
    |-- greetings_test.go
```

and here is the code of greetings.go
```go
package greetings

import (
	"fmt"
)

func Hello() {
    fmt.Println("Hello")
}
```

## Example of how to call the code from different module
Assume that the greetings module is published to the public repository. To call the `Hello` function in the greetings module, we need to import it into the code.
```
import (
	"example.com/greetings"
)

func main() {
    greetings.Hello()
}
```