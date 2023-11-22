package main

import "fmt"

type Book struct {
	name string
}

func (book *Book) SetName(name string) {
	book.name = name
}
func main() {
	book := &Book{name: "facebook"}
	book.SetName("hello")
	fmt.Println(book.name)

}
