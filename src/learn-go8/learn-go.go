package main

import (
	"fmt"
	"sync"
)

//匿名结构体：struct
func testStruct() {
	//匿名结构体的语法
	// struct{
	// 	property type
	// 	property type
	// }{
	// 	property:val,
	// 	property:val,
	// }
	all := []struct {
		name string
		age  int
	}{
		{
			name: "kiwi",
			age:  18,
		},
		{
			name: "apache",
			age:  22,
		},
		{
			name: "java",
			age:  95,
		},
	}
	for _, val := range all {
		fmt.Println(val)
	}
}

// Go语言在select语句中实现优先级
//select 会阻塞，直到其中的一个 channel 转为就绪状态时执行对应case分支的代码。
//select本质就是底层select多路复用器
//当select 存在多个 case时会随机选择一个满足条件的case执行。
func testSelect(ch []chan string) {
	for _, val := range ch {
		select {

		case node := <-val:
			fmt.Println("---->", node)
		}
	}
}

func v() {
	var wg sync.WaitGroup = sync.WaitGroup{}

	wg.Add(2)
	ch := []chan string{
		//无缓存，会被阻塞住，但是使用了 go这个关键字，后台开启协程，发送阻塞的代码会被另一个协程go接受
		make(chan string),
		make(chan string),
		make(chan string),
	}
	go func() {
		for _, val := range ch {
			fmt.Println(val)
			val <- "🐏"

		}
		wg.Done()
	}()
	go func() {
		testSelect(ch)
		wg.Done()
	}()
	wg.Wait()
}

func buffer() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case x := <-ch:
			fmt.Println(x)
		}
	}
}

func testcase() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 1
	select {
	case i := <-ch2:
		fmt.Println("===", i)
	case i := <-ch1:
		fmt.Println("---", i)

	}
}
func main() {
	v()

}
