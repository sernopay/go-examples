# Implementing the Pillars of OOP in Go
As we know, Go is not a pure OOP language, but it does support OOP principles. In this section, we will explore how to implement the four pillars of OOP in Go: Abstraction, Inheritance, Encapsulation, and Polymorphism. We are going to use a simple example of a Document Printer that can print the content of different types of documents (e.g., PDF, Word, etc.) to illustrate how to implement these principles in Go.

## Learning Objectives
- Understand how to implement the four pillars of OOP in Go.

## Abstraction
Abstraction is the process of hiding the implementation details and showing only the essential features of the object. We can achieve abstraction in Go by using interfaces and struct types. An interface defines a set of methods that a concrete struct must implement, while a base struct defines the general behavior of the object. \
In this example, we will create an interface called `Printer` that defines the behavior of a document printer. Also we will create an "abstract class" called `DocumentPrinter` that will define the general behavior of a document printer. Later, we will create concrete implementation of this object for different type of documents (e.g., PDF, Word, etc.) in the inheritance section.

```go
package main

import "fmt"

// DocumentPrinter is the "abstract class" that defines the general behavior of a document printer.
// This struct should not be instantiated directly, but should be "extended" by concrete types.
type DocumentPrinter struct {
	content string
}

func (d *DocumentPrinter) Print() {
	fmt.Println(d.content)
}

// The interface defines the behaviour of the object without specifying how it is implemented.
// We can think of it as an abstract method that must be implemented by the concrete types.
type Printer interface {
	Read()
	Print()
}
```

From the code above, we can see that we have defined a "base struct" called `DocumentPrinter` as an "abstract class". This struct has a generic method called `Print`, which will be shared accross all the concrete types. \
We also defined an interface called `Printer`, which defines the behavior of the object without specifying how it is implemented. Some methods, such as `Print`, are already implemented in the base struct, while others must be implemented by the concrete types. The unimplemented methods can be thought of as abstract methods, similiar to those in other OOP languages.

## Inheritance
Inheritance is the process of creating a new class from an existing class. The new class inherits the properties and behaviors of the existing class. In Go, inheritance is achieved through composition by embedding structs. This allows the embedded struct's fields and methods to become part of the outer struct. \
Continuing with our example, we will create two concrete types called `PDFPrinter` and `WordPrinter` that will "inherit" from the `DocumentPrinter` struct and implement the `Printer` interface.

```go
// WordPrinter
type WordPrinter struct {
	*DocumentPrinter // Embedding the DocumentPrinter struct
}

func (w *WordPrinter) Read() {
	w.content = "Word Converted Content"
}

// PDFPrinter
type PDFPrinter struct {
	*DocumentPrinter // Embedding the DocumentPrinter struct
}

func (p *PDFPrinter) Read() {
	p.content = "PDF Converted Content"
}

func (p *PDFPrinter) Print() {
	fmt.Printf("Print PDF: %s \n", p.content)
}
```

In the code above, we have created two concrete types called `WordPrinter` and `PDFPrinter` that "inherit" from the `DocumentPrinter` struct. We have also implemented the `Read` method for both types. The `Print` method is already implemented in the base struct, so we don't need to implement it again. However, the Print method in `PDFPrinter` shadows the Print method of the embedded `documentPrinter`, providing a custom implementation.

## Encapsulation
Encapsulation is the process of restricting access to an object's internal state and allowing access only through public methods. In Go, this is achieved using exported and unexported fields and methods. Exported fields and methods are accessible from outside the package, while unexported ones are only accessible within the same package. \
In our example, we will unexport the `DocumentPrinter` struct since we don't want to allow instantiation of this struct directly. We will also unexport the other concrete types, as their initialization requires initializing the base struct as well. Since Go does not support constructors, 
we will create a "constructor" function to initialize the concrete type and the base struct.

```go
// We change the struct name from `DocumentPrinter` to `documentPrinter` to make it unexported. 
type documentPrinter struct {
	content string
}

func (d *documentPrinter) Print() {
	fmt.Println(d.content)
}

// We also change the struct name of the concrete types to make them unexported.
type wordPrinter struct {
	*documentPrinter
}

func (w *wordPrinter) Read() {
	w.content = "Word Converted Content"
}

type pdfPrinter struct {
	*documentPrinter
}

func (p *pdfPrinter) Read() {
	p.content = "PDF Converted Content"
}

func (p *pdfPrinter) Print() {
	fmt.Printf("Print PDF: %s \n", p.content)
}
```

After making the struct unexported (by starting its name with a lowercase letter), it cannot be instantiated directly from outside the package.

## Polymorphism
Polymorphism is the ability of different types to be treated as the same type through a common interface. In Go, we can achieve polymorphism by using interfaces. In our case, we have already defined an interface called `Printer` that defines the behavior of a document printer. We can use this interface to create a function that accepts or returns any type that implements the `Printer` interface.

```go
type Printer interface {
	Read()
	Print()
}

type documentPrinter struct {
	content string
}

...

func NewWordPrinter() Printer {
	return &wordPrinter{
		documentPrinter: &documentPrinter{},
	}
}

func NewPDFPrinter() Printer {
	return &pdfPrinter{
		documentPrinter: &documentPrinter{},
	}
}
```

In the code above, we have created functions called `NewWordPrinter` and `NewPDFPrinter` that return a `Printer` interface. Since both `wordPrinter` and `pdfPrinter` implement the `Printer` interface, we can use them interchangeably. This is an example of polymorphism in Go.

## Run the Code
To run the code, we need to create a main function that will create instances of the concrete types and call their methods.
```go
func main() {
	wordP := NewWordPrinter()
	wordP.Read()
	wordP.Print()

	pdfP := NewPDFPrinter()
	pdfP.Read()
	pdfP.Print()
}
```

The output of the code will be:
```bash
Word Converted Content
Print PDF: PDF Converted Content
```
The output demonstrates polymorphism, where both `wordPrinter` and `pdfPrinter` implement the `Printer` interface and provide their own implementations of the Read and Print methods.

## Summary
In this section, we have explored how to implement the four pillars of OOP in Go: Abstraction, Inheritance, Encapsulation, and Polymorphism. We have used a simple example of a Document Printer that can print the content of different types of documents (e.g., PDF, Word, etc.) to illustrate how to implement these principles in Go. We have also explored how interfaces and struct types enable abstraction and polymorphism, and how embedding allows us to achieve inheritance. Finally, we have seen how to use exported and unexported fields and methods to achieve encapsulation.