package main

import (
	"fmt"
)

//传递一个接口引用	
func add(a int, b int ,c compartor) int{
	return c.compartor(a,b)
}

type compartor interface{
	compartor(a int,b int) int
}

//用结构体A去实现compartor方法，进而实现compartor接口
type A struct{}
func (a A)compartor(x int,y int) int{
	return x-y
}

//空接口的使用,可以保存任何数据类型（Object超类）
//空接口的应用：
//	1.空接口类型作为函数的参数  （比如fmt.Println()这个函数）
//	2.空接口可以作为map的value   make(map(string)interface{})  这样就可以 map["name"] = "kiwi"  map["age"] = 18
func EmptyInterface(){
	var a interface{}
	a = "hello"
	fmt.Println(a)
	a = false
	fmt.Println(a)
	a = 15
	fmt.Println(a)
}

//类型断言
//	语法：x.(T)
//	return: val,ok
//测试code: fmt.Println(test())
func test() interface{}{
	var a interface{}
	a = "string"
	a = 18
	val,ok := a.(string)
	if !ok{
		fmt.Println("类型断言错误")
		return 0
	}
	fmt.Println(val)
	return "ok"
}

//类型断言 switch
func test2(a interface{}){
	//type关键字   a.(type)返回一个类型变量v
	switch v := a.(type) {
		case string:
			fmt.Println("字符串类型",v)
		case bool:
			fmt.Println("布尔类型",v)
		case  int:
			fmt.Println("int类型",v)
		default:
			fmt.Println("拆不到",v)		
	}
	
}

func main() {
	// EmptyInterface()
	// fmt.Println(add(15,23,A{}))
	// fmt.Println(test())
	fmt.Println()
}