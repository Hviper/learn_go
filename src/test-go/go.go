package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//1 秒后全部协程结束
	defer cancel()

	go handle(ctx, 900*time.Millisecond)
	select {
	// 多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，
	//一旦接收到取消信号就立刻停止当前正在执行的工作。
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	// 多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，
	//一旦接收到取消信号就立刻停止当前正在执行的工作。
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
