package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := make([]int, 0, 8)
	Append(a)
	(*reflect.SliceHeader)(unsafe.Pointer(&a)).Len = 1
	fmt.Println(a)
}

func Append(a []int) {

	a = append(a, 1, 2, 3)

}

func Change(tmpb []int) {
	tmpb[0] = 10
}
