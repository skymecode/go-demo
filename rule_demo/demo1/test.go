package main

import "fmt"

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

// Write
func (s *S) Write(content string) {
	s.data = content
}
func (s S) WriteOther(str string) {
	s.data = str
	fmt.Println("值传递", s.data)
}

var sValues = map[int]S{1: {"A"}}

func main() {
	read := sValues[1].Read()
	fmt.Println(read)
	//sValues[1].Write("hello")
	sParts := map[int]*S{1: {"A"}}
	sParts[1].Write("hello")
	read = sParts[1].Read()
	fmt.Println(read)
	sValues[1].WriteOther("test")
	s := sValues[1].Read()
	fmt.Println(s)
}
