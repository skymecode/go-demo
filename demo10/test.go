package main

import (
	"fmt"
)

type field struct {
	num int
}

func (t *field) print(n int) {
	fmt.Println(t.num, n)
}
func main() {
	var i int = 3000000000000000000
	defer fmt.Println("result2 =>", func() int { return i * 2 }()) //声明的时候已经确认i==1了
	i++

	v := field{1}
	defer v.print(func() int { return i * 2 }()) //确定了i==2,但是再回来调用的时候的v是最后确定的传递过来的
	v = field{2}
	i++

	// prints:
	// 2 4
	// result => 2 (not ok if you expected 4)
}
