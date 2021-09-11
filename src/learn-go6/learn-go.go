package main
//IO操作
import (
	"fmt"
	"os"
	"io/ioutil"
)


// var i MyInterface
// i = tools{
// 	name:"何东昌",
// }
// fmt.Println(i.Name())
// fmt.Println(i.Size())
type  tools struct{
	name string
}
type MyInterface interface{
	Name() string  
	Size() int
}
func (t tools)Name()string{
	return "kiwi"
}
func (t tools) Size()int{
	return 10086
}


//"io/ioutil"  ---->ioutil
func readByioutil(){
	content,err := ioutil.ReadFile("E:/go/src/learn-go6/index.txt")
	if err!=nil{
		fmt.Println("错误",err)
		return 
	}
	//读取出来的是content为字节数组，转为字符串
	fmt.Println(string(content))
}

//os
func readFile(){
	fileObj,err := os.Open("E:/go/src/learn-go6/index.txt")
	if err!=nil{
		fmt.Println("错误！！",err)
		return
	}
	var tempCap []byte = make([]byte,1024)
	for{
		fn,err := fileObj.Read(tempCap)
		if err!=nil{
			fmt.Println("错误xx")
			break
		}
		fmt.Println(fn)
		fmt.Println(string(tempCap))
	}
	defer fileObj.Close()
}
func main() {
	readByioutil()
	

}