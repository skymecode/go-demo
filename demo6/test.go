package main

import (
	"fmt"
	"time"
)

// 定义go协程池结构体
type GoPool struct {
	MaxLimit  int           // 最大并发度
	tokenChan chan struct{} // 控制并发执行的令牌通道
}

func main() {
	gopool := NewGoPool(3) // 最大并发度为 3
	defer gopool.Wait()
	//开启10个线程，但最大并发度为3
	for i := 0; i < 10; i++ {
		taskID := i
		gopool.Submit(func() {
			// 模拟任务处理逻辑
			fmt.Printf("任务 %d 开始执行\n", taskID)
			time.Sleep(time.Second)
			fmt.Printf("任务 %d 完成\n", taskID)
		})
	}
}

// NewGoPool 创建一个 GoPool 对象，max 设置最大并行度
func NewGoPool(max int) *GoPool {
	p := &GoPool{}
	p.MaxLimit = max
	p.tokenChan = make(chan struct{}, p.MaxLimit)
	for i := 0; i < p.MaxLimit; i++ {
		p.tokenChan <- struct{}{} // 初始化 tokenChan
	}
	return p
}

// Submit 提交任务
// 每个任务都会异步执行，并且会在执行完后释放一个令牌
func (gp *GoPool) Submit(fn func()) {
	token := <-gp.tokenChan // 如果没有可用令牌，则会阻塞

	go func() {
		fn()
		gp.tokenChan <- token // 执行完成后释放令牌，以便其他任务可以获得执行机会
	}()
}

// Wait 等待所有任务执行完成
func (gp *GoPool) Wait() {
	for i := 0; i < gp.MaxLimit; i++ {
		<-gp.tokenChan // 等待所有令牌被释放
	}

	close(gp.tokenChan) // 关闭令牌通道，释放资源
}

func (gp *GoPool) Size() int {
	return len(gp.tokenChan) // 返回当前令牌通道中的可用令牌数量
}
