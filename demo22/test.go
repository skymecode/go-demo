package main

import "fmt"

func main() {

	defer fmt.Println("main...结束")
	fmt.Println("main....运行")
	r := make(chan int)
	go func(i int) {
		defer fmt.Println("g-r结束")
		r <- i
	}(5)
	num := <-r
	fmt.Println(num)
}
