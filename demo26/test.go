package main

import "fmt"

func main() {
	reverseWords("  hello world  ")
}
func reverseWords(s string) string {
	b := []byte(s)
	fmt.Println(&b)
	removeElement(b)
	//这里又采坑了,len没变！
	fmt.Println("a", &b)
	fmt.Println(string(b), "j")
	reverse(b)
	last := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			left, right := last, i-1
			for left < right {
				b[left], b[right] = b[right], b[left]
				left++
				right--
			}

			last = i + 1
		}
	}
	ans := string(b)
	return ans
}
func removeElement(b []byte) {
	//双指针法->快慢指针,复杂度为o(n)
	slow := 0
	for i := 0; i < len(b); i++ {
		if b[i] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for i < len(b) && b[i] != ' ' { // 复制逻辑
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	fmt.Println(&b)
	b = b[0:slow]
	fmt.Println("new", &b)

}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
