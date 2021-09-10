package main


import (
	"fmt"
)

type cat struct{
	name string
	age int 
}
//相当于继承了MyCat
type MyCat cat
func (c MyCat)test(s string) (string,int){
	fmt.Println(c)
	fmt.Println(s)
	return "hello",18
}
func main(){
	c := MyCat{
		name:"hello",
		age:18,

	}
	fmt.Println(c.test("ppppp"))
}