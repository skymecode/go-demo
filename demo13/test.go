package main

import (
	"fmt"
	"reflect"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}
type Book struct {
}

func (book *Book) WriteBook() {
	fmt.Println("写操作")
}

func (book *Book) ReadBook() {
	fmt.Println("读操作")
}

func main() {
	book := &Book{}
	book.ReadBook()

	var r Reader
	r = book
	fmt.Printf("类型:%T\n", r)
	r.ReadBook()

	var w Writer

	of1 := reflect.TypeOf(w)
	fmt.Println(of1)
	w = r.(Writer)
	w.WriteBook()

	of2 := reflect.TypeOf(w)
	fmt.Println(of2)
}
