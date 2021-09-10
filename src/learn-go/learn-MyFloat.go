package main
import (
	"fmt"
)
type MyFloat float64

//拓展方法
func (n MyFloat) show() (string){
	return "myFun"
}
func main(){
	var n MyFloat = 15.2
	fmt.Println(n.show())
}