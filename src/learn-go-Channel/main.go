package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct{}

func (t task) Run() error {
	return myError{}
}

type myError struct{}

func (e myError) Error() string {
	return "错误了"
}

func test() {
	tasks := []task{
		task{},
		task{},
		task{},
		task{},
		task{},
		task{},
	}

	errCh := make(chan error, len(tasks)) //缓存队列
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	for i, _ := range tasks {
		if i == (len(tasks) - 1) {
			fmt.Println("全部任务执行完成了。。。", i)
		}
		go func() {
			defer wg.Done()
			// if条件表达式的前面可以包含初始化语句，支持平行赋值，但不支持多个赋值语句
			// 赋值+条件判断
			if err := tasks[i].Run(); err != nil {
				errCh <- err
			}

		}()
	}
	fmt.Println("等待。。。")
	wg.Wait()
	fmt.Println("开始循环下面代码。。。")
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case err := <-errCh:
			fmt.Println("接受信号", err)
		default:
			fmt.Println("等待发送信号")
		}
	}
}
func main() {
	test()
}
