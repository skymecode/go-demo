package main

import "fmt"

func main() {
	//lib1.Test()
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
	var p *int
	p = &a
	var pp **int
	pp = &p //指向指针的地址的指针
	*pp = &b
	fmt.Println(*p)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
