package main

import "fmt"

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}

type Book struct {

}

func (this *Book) ReadBook(){
	fmt.Println("Read Book")
}

func (this *Book) WriteBook()  {
	fmt.Println("Write Book")
}

func main() {
	b := &Book{}

	var r Reader
	r= b

	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()
}