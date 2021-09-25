package main

import (
	"fmt"
	"runtime"
)

//对于 go 中的方法，当接收者是值类型的时候，调用者可以是值类型，也可以是指针引用类型

func test1() {
	for i := 0; i < 5; i++ {
		fmt.Println("11111---------------->", i)
	}
}

func test2() {
	for i := 0; i < 5; i++ {
		fmt.Println("22222---------------->", i)
	}
}
func main() {
	go test1()
	go test2()

	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("main run")
	}
}
