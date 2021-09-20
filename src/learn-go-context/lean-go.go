package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func groupSync() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		target(5, "🐏")
		wg.Done()
	}()
	go func() {
		target(5, "🐂")
		wg.Done()
	}()
	wg.Wait()
}

func target(count int, s string) {
	for i := 0; i < count; i++ {
		fmt.Println(s)
	}
}

func contextFunc() {

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "  ---->协程任务被迫结束")
				return
			default:
				fmt.Println(name, " working")
			}
		}
	}(ctx, "go11")
	go func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "  ---->协程任务被迫结束")
				return
			default:
				fmt.Println(name, " working")
			}
		}
	}(ctx, "go22")

	go func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "  ---->协程任务被迫结束")
				return
			default:
				fmt.Println(name, " working")
			}
		}
	}(ctx, "go33")

	//设置主线程最大等待时间
	time.Sleep(50 * time.Millisecond)
	//当执行到cancel()的时候，即私有协程对象全部执行结束，不会再等待，防止超时处理
	cancel()
	time.Sleep(5 * time.Millisecond)
}

func main() {
	contextFunc()
}
