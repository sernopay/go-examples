# Calling a Local Module Function
This usually happens when you are still developing two modules at the same time, and neither module has been published to the repository yet.

In this example, the hello module wants to call a function from the greetings module. To do that use the `go mod edit` command.

```
go mod edit -replace example.com/greetings=../greetings
```

Assuming it has the following folder structure:
```
<home>/
 |-- greetings/
 |-- hello/
```

After you run the command, the go.mod file in the hello directory should include a replace directive
```
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings
```

run the `go mod tidy` command to synchronize the example.com/hello module's dependencies
