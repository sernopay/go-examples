package main

import "fmt"

// Abstraction

// documentPrinter is the "abstract class" that defines the general behavior of a document printer.
// This struct should not be instantiated directly, but should be "extended" by concrete types.
type documentPrinter struct {
	content string
}

func (d *documentPrinter) Print() {
	fmt.Println(d.content)
}

// The interface defines the behaviour of the object without specifying how it is implemented.
// We can think of it as an abstract method that must be implemented by the concrete types.
type Printer interface {
	Read()
	Print()
}

// Inheritance

// WordPrinter
type wordPrinter struct {
	*documentPrinter
}

func (w *wordPrinter) Read() {
	w.content = "Word Converted Content"
}

// PDFPrinter
type pdfPrinter struct {
	*documentPrinter
}

func (p *pdfPrinter) Read() {
	p.content = "PDF Converted Content"
}

func (p *pdfPrinter) Print() {
	fmt.Printf("Print PDF: %s \n", p.content)
}

// Constructor functions
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

// main
func main() {
	wordP := NewWordPrinter()
	wordP.Read()
	wordP.Print()

	pdfP := NewPDFPrinter()
	pdfP.Read()
	pdfP.Print()
}
