package main

import (
	"fmt"
	"time"
)

const (
	i = 1 << iota
	j = 3 << iota
	k = 1 << iota
	l
)

func main() {
	//fmt.Println("i=", i)
	//fmt.Println("j=", j)
	//fmt.Println("k=", k)
	//fmt.Println("l=", l)
	//var a int
	//var ptr *int
	//a = 1
	//if a > 0 {
	//	fmt.Printf("ptr的变量类型%T\n", ptr)
	//	ptr = &a
	//	fmt.Printf("%d", *ptr)
	//} else {
	//	fmt.Println("a不大于0")
	//}
	//scan, err := fmt.Scan(&a)
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(scan)
	//if a == 123456 {
	//	fmt.Println("密码正确")
	//} else {
	//	fmt.Println("密码错误")
	//}
	//switch {
	//case a == 1:
	//	fmt.Println("a==1")
	//case a == 2:
	//	fmt.Println("a==2")
	//default:
	//	{
	//		fmt.Println("hello")
	//		fmt.Println("world")
	//	}
	//}
	//var x interface{}
	//switch i := x.(type) {
	//case nil:
	//	fmt.Printf("x的类型是%T", i)
	//case int:
	//	fmt.Printf("x是int%T", i)
	//default:
	//	fmt.Printf("未知类型")
	//
	//}
	//fallthrough 只会影响紧随其后的下一个 case 分支，而不会影响之后的分支。
	//
	//fallthrough 不会检查 case 分支的条件，它只是简单地继续执行下一个 case 分支，而不管条件是否匹配。
	//
	//使用 fallthrough 应该小心谨慎，因为它可能会导致不常见的逻辑错误。通常情况下，fallthrough 并不是必要的，因为 switch 语句会自动终止，除非明确使用 fallthrough。
	//
	//大多数情况下，你可以使用 case 分支来进行条件匹配，而不需要使用 fallthrough。只有在某些特定的情况下才需要使用 fallthrough 来达到穿透的目的。
	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)
	go Chann(ch, stopCh)
	for {
		select {
		case c = <-ch:
			fmt.Println("Receive C", c)
		case s := <-ch:
			fmt.Println("Receive D", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}
func Chann(ch chan int, stopCh chan bool) {
	for j := 0; j < 10; j++ {
		ch <- j
		fmt.Printf("函数执行%d\n", j)
		time.Sleep(time.Second)
	}
	stopCh <- true
}
