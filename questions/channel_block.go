package questions

import (
	"context"
	"fmt"
	"time"
)

/*
题目 1:程序退出时是否存在 goroutine 泄漏？
为什么？
如何修复？
*/
func worker(ch chan int) {
	for v := range ch {
		fmt.Println("got:", v)
	}
}

func RunWorker() {
	ch := make(chan int)

	go worker(ch)

	// 发送 3 次后不再发送
	ch <- 1
	ch <- 2
	ch <- 3
}

/*
题目 2：这个 select 超时是否一定可靠？
问题：
会不会 timeout 永远不触发？
如果 ch 一直阻塞，timeout 能否打断 case v := <-ch？
如何保证 timeout 一定生效？
如果 ch 数据特别密集，会不会永远不 timeout？
*/
func listen(ch <-chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println("value:", v)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
			return
		}
	}
}

func listenFixed(ctx context.Context,ch <-chan int) {
	for {
		select {
			case <-ctx.Done():
				fmt.Println("context done")
				return
		case v,ok := <-ch:
			if !ok{
				fmt.Println("channel closed")
			}
			fmt.Println("value:", v)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
			return
		}
	}
}
/*
题目 3：为什么这个 select 会卡死？
为什么主 goroutine 卡住？
process() 退出了吗？
ch 是否泄漏？
应该如何正确优雅退出？
*/
func process(ch chan int, done chan struct{}) {
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			fmt.Println("exit")
			return
		}
	}
}

func RunProcess() {
	ch := make(chan int)
	done := make(chan struct{})

	go process(ch, done)

	close(done)
}

