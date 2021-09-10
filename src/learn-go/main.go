package main

import (
	"fmt"
)


func delp() int {
	
	return 1
}

// func function_name(input1 type1, input2 type2) (type1, type2) {
// 	函数体
// 	返回多个值
// 	return value1, value2
//  }

func add(a int,b int) int{
	numbers := make(map[string]int)
	numbers["name"] = 15
	numbers["age"] = 15
	fmt.Println(numbers["name"])
	delete(numbers,"name")
	fmt.Println("这是水水水水水水-->",numbers)
	return a+b
}

func apache(a *int){
	*a +=1
}
func main(){
	arr := [] string{"aaa","bbb"}
	test(arr)
	fmt.Println(arr)

}
func forEach(arr [] int){
	for k,v:=range arr{
		fmt.Println(k,v)
	}
}

func test(arr [] string) ([] string){
	arr = append(arr,"hahha")
	return arr
}



