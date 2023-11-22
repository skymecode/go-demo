package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func doSomething() error {
	originalErr := errors.New("An original error occurred")
	return errors.Wrap(originalErr, "Error in doSomething")
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Full error:", err)
		originalErr := errors.Cause(err)
		fmt.Println("Original error:", originalErr)
	}
}
