package main

import "fmt"

type errNotFound struct {
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func open(file string) error {
	return errNotFound{file: file}
}

func use() {
	if err := open("hello"); err != nil {
		if _, ok := err.(errNotFound); ok { //类型断言,判断是否有指定类型的值
			// handle
			fmt.Println("ok")
		} else {
			panic("unknown error")
		}
	}
}
func main() {
	use()
}
