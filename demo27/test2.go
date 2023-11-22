package main

import (
	"fmt"
	"time"
)

func test2() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("协程结束")
		defer close(c)
		for i := 0; i < 3; i++ {
			c <- i
		}

	}()
	time.Sleep(2 * time.Second)
	go func() {
		defer fmt.Println("读取结束")

	}()

	for {
		if num, ok := <-c; ok {
			fmt.Println(num)
		} else {
			break
		}

	}
}
