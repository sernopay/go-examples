# Fuzzing

With fuzzing, random data is run against your test in an attempt to find vulnerabilities or crash-causing inputs. Some examples of vulnerabilities that can be found by fuzzing are :
* SQL injection
* buffer overflow
* denial of service
* cross-site scripting attacks

One benefit of fuzzing is it comes up with inputs for your code and may identify edge cases that the test cases you came up with didn’t reach.

Fuzzing has a few limitations as well. 
* In your unit test, you could predict the expected output of the Reverse function and verify that the actual output met those expectations. For example  in the test case `Reverse("Hello, world")` the unit test specifies the return as `"dlrow ,olleH"`.
* When fuzzing, you can’t predict the expected output, since you don’t have control over the inputs.

the syntax differences between the unit test and the fuzz test:
* The function begins with FuzzXxx instead of TestXxx and takes *testing.F instead of *testing.T
* Where you would expect to see a t.Run execution, you instead see f.Fuzz which takes a fuzz target function whose parameters are *testing.T and the types to be fuzzed.

This Go program demonstrates fuzz testing using the `testing` package. Fuzz testing is a type of automated testing that provides random data as input to a function to find bugs and unexpected behavior.

## Prerequisites
An installation of Go 1.18 or later

## Usage

To run the fuzz test, use the following command:
```
go test
```
This will start the fuzzing process and test the Reverse function with various inputs to find potential bugs.

You can also run go test -run=FuzzReverse if you have other tests in that file and you only wish to run the fuzz test.
```
go test -fuzz=FuzzReverse
```

Another useful flag is -fuzztime which restricts the time fuzzing takes. For example specifying `-fuzztime 10s` in the test  would mean that, as long as no failures occurred earlier, the test would exit by default after 10 seconds had elapsed

To run a specific corpus entry within FuzzXxx/testdata  you can provide {FuzzTestName}/{filename} to -run. This can be helpful when debugging.
```
go test -run=FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0
```

## To see if any randomly generated string inputs will cause a failure
the input that caused the problem is written to a seed corpus file that will be run the next time go test is called even without the -fuzz flag. \
To view the input that caused the failure open the corpus file written to the testdata/fuzz/FuzzReverse directory in a text editor Your seed corpus file may contain a different string, but the format will be the same.
```
go test fuzz v1
string("泃")
```
The first line of the corpus file indicates the encoding version. Each following line represents the value of each type making up the corpus entry. Since the fuzz target only takes 1 input, there is only 1 value after the version.