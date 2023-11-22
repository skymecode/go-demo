package main

import (
	"errors"
	"fmt"
	"math"
)

type Phone struct {
	num string
}

func main() {
	iPhone := Phone{
		num: "15881330597",
	}
	var s *Phone
	s = &iPhone
	fmt.Println(s.num)
	nokia := NewNokiaPhone("Nokia 123")
	fmt.Println(nokia.model)
	_, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	}
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	sqrt := math.Sqrt(f)
	return sqrt, nil
}

type NokiaPhone struct {
	model string
}

func NewNokiaPhone(model string) *NokiaPhone {
	return &NokiaPhone{
		model: model,
	}
}
