package main
import (
	"fmt"
	
)

//切片是一个引用类型
func divide(arr []int) []int{
	a := append(arr,15)
	return a
}	


//函数传递的都是值传递，即使是指针传递也是传递的是指针的一个副本，修改副本对原指针没有影响
func P(a *[] int){
	fmt.Printf("%p\n",*a)
	a = &[] int{1,2,3}
	
}



/*
append会发生值拷贝，如果是自身对象会复用，所以地址不会变化，如果你赋值给另一个变量则直接完全拷贝一份内存给这个变量进行引用
append赋值给一个新对象时（新变量），因为是值拷贝，所以会是一个新对象
切片翻倍扩容时，ptr会发生数据拷贝，编程一个新指针
append之后的新对象因为拷贝了ptr指针，所以ptr对象不会变化
append赋值前后ptr
*/
// var a []int
// for i :=0;i<10;i++{
// 	a = append(a,i)    //最终赋值给自身，复用
// 	fmt.Printf("%v len:%d cap:%d ptr:%p \n",a,len(a),cap(a),a)
// }
//result：
// 	[0] len:1 cap:1 ptr:0xc00000a0a8 
// 	[0 1] len:2 cap:2 ptr:0xc00000a0f0 
// 	[0 1 2] len:3 cap:4 ptr:0xc0000122a0 
// 	[0 1 2 3] len:4 cap:4 ptr:0xc0000122a0 
// 	[0 1 2 3 4] len:5 cap:8 ptr:0xc00000e280 
// 	[0 1 2 3 4 5] len:6 cap:8 ptr:0xc00000e280 
// 	[0 1 2 3 4 5 6] len:7 cap:8 ptr:0xc00000e280 
//	[0 1 2 3 4 5 6 7] len:8 cap:8 ptr:0xc00000e280 
// 	[0 1 2 3 4 5 6 7 8] len:9 cap:16 ptr:0xc00001c100 
// 	[0 1 2 3 4 5 6 7 8 9] len:10 cap:16 ptr:0xc00001c100 


//数组追加数组：   ...展开运算符语法糖
// arr := []int {1,2,3,4,5,6}
// array := make([]int,0)
// array = append(array,arr...)



/*
结构体的可见性。结构体中字段大写开头表示可公开访问，小写表示私有
（私有：仅在定义当前结构体的包中可以访问）
如果一个Go语言包中定义的标识符是首字母大写的，那么就是对外可见的
*/


func main(){
	
	arr := []int {1,2,3,4,5}
	fmt.Println(arr)
}


