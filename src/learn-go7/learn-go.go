package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func reflectType(x interface{}) {
	//程序开始我是不知道别人调用我这个函数传递进来是什么类型(java中的getClass)
	//1.方法1：通过类型断言
	//2.方法2：借助反射 【运行时动态的获取对应值的类型】
	obj := reflect.TypeOf(x)
	fmt.Println(obj)
}

//多路复用器
//Don't communicate by sharing memory;
//share memory by communicating
//chan T   // 可以接收和发送类型为 T 的数据
func selectTest() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "🐏"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "🐂"
			time.Sleep(time.Millisecond * 200)
		}
	}()
	for {
		//select epoll 多路复用器
		//select保证可以在第一时间从新消息的channel里收到数据
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}

	}
}

//发起者手动关闭chan
func count(n int, c chan string) {
	for i := 0; i < n; i++ {
		c <- "🐏"
	}
	//关闭chan
	close(c)
}

//Buffered Channels
//通过缓存的使用，可以尽量避免阻塞，提供应用的性能。
func BufferChanTest() {
	// var wg sync.WaitGroup
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int, 1)
	go func() {
		ch <- 15
		ch <- 15
		ch <- 15
		ch <- 15
		wg.Done()
	}()
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}

func main() {

	BufferChanTest()

}
