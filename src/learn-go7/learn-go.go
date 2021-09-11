package main

import(
	"fmt"
	"reflect"
)

func reflectType(x interface{}){
	//程序开始我是不知道别人调用我这个函数传递进来是什么类型(java中的getClass)
	//1.方法1：通过类型断言
	//2.方法2：借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj)	
}

func main() {
	reflectType("stio")
	reflectType(12)
}