package main

import "fmt"

func main() {
	s := make([]string, 0)
	empty := isEmpty(s)
	fmt.Println(s)
	fmt.Println(empty)
}
func isEmpty(s []string) bool {
	return len(s) == 0
}
