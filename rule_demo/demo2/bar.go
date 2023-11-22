package main

import foo "GoProject/rule_demo/demo3"

func main() {
	if err := foo.Open(); err != nil {
		if err == foo.ErrCouldOpen {
			// handle
		} else {
			panic("unknown error")
		}
	}
}
