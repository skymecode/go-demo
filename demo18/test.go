package main

import (
	"fmt"
)

func main() {
	var a = []byte{98, 99, 100}
	s := string(a)
	fmt.Println(s)
	bytes := []byte(s)
	bytes[0] = 'A'
	s2 := string(bytes)
	fmt.Println(s2)
	sprintf := fmt.Sprintf("hello%v", s2)
	fmt.Println(sprintf)
}
