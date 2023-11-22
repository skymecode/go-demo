package main

import (
	"fmt"
	"time"
)

func main() {
	selectTest()
}
func test() {
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
func fib(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			{
				x, y = y, x+y

			}
		case <-quit:
			{
				fmt.Println("退出")
				return
			}
		}
	}

}
func selectTest() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fib(c, quit)

}
