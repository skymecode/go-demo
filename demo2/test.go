package main

import (
	"fmt"
	"math"
)

func main() {
	//strings := []string{"google", "runoob"}
	//for i, s := range strings {
	//	fmt.Println(i, s)
	//}
	//numbers := [6]int{1, 2, 3, 5}
	//for i, x := range numbers {
	//	fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	//}
	//i := 10
	//for i > 0 {
	//	fmt.Printf("%d", i)
	//	i--
	//}
	var a int = 100
	var b int = 200
	swap(&a, &b)
	fmt.Println(a, b)
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(10.001))
	add := func(a, b int) int {
		return a + b
	}
	sub := func(a, b int) int {
		return a - b
	}

	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 2, 8)
	sum2 := calculate(sub, 10, 1)
	fmt.Println(sum2)
	fmt.Println("2 + 8 =", sum)
	difference := calculate(func(a int, b int) int {
		return a + b
	}, 10, 4)
	fmt.Println(difference)
}

// 地址给到指针,那么指针的值变化也就是真正意义上的值的变化
func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
