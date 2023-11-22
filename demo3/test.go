package main

import "fmt"

func main() {
	var numbers []int

	/* 允许追加空切片 */
	numbers = append(numbers, 0)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3)
	printSlice(&numbers)
	fmt.Println(numbers)

	///* 创建切片 numbers1 是之前切片的两倍容量*/
	//numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	//
	///* 拷贝 numbers 的内容到 numbers1 */
	//copy(numbers1, numbers)
	//printSlice(numbers1)
	m := make(map[string]string)
	m["hello"] = "world"
	m["Google"] = "搜索"
	for key, value := range m {
		fmt.Printf("key=%s和value=%s\n", key, value)
	}
	var a int32
	a = 17
	var mean float64
	mean = float64(a) / float64(5)
	fmt.Printf("%.3f", mean)

}

func printSlice(x *[]int) {
	*x = append(*x, 666)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(*x), cap(*x), x)
}
