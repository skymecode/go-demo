package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 实现heap.Interface的Less方法
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// 实现heap.Interface的Swap方法
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 实现heap.Interface的Push方法
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 实现heap.Interface的Pop方法
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func heapSort(nums []int) []int {
	var h IntHeap
	heap.Init(&h)

	// 将切片元素依次推入堆中
	for _, num := range nums {
		heap.Push(&h, num)
	}

	// 依次弹出堆顶元素，即为升序排列的结果
	sorted := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sorted[i] = heap.Pop(&h).(int)
	}

	return sorted
}

func main() {
	isHappy(2)
	input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("未排序前:", input)

	// 调用堆排序函数
	output := heapSort(input)

	fmt.Println("排序后:", output)

}
func isHappy(n int) bool {
	//但模拟会出现最终边界模糊的问题
	//如果这个结果重复出现那就退出！
	m := make(map[int]bool)
	var sum int
	var s = n
	for s != 1 && !m[n] {
		temp := n % 10
		sum = sum + (temp * temp)
		n = n / 10

		if n == 0 {
			if sum == 1 {
				return true
			}
			if !m[s] {
				m[s] = true
			} else {
				return false
			}
			n = sum
			s = n
			sum = 0
		}
	}
	return s == 1
}
