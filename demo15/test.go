package main

import "fmt"

func main() {
	var vstat [3]int
	vstat[0] = 1
	vstat[1] = 2
	vstat[2] = 3
	var vauto *[3]int = new([3]int)
	*vauto = vstat
	fmt.Println(vauto)
}
