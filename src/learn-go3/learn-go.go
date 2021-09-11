package main

import(
	"fmt"
	"encoding/json"
)
//二进制  八进制   十六进制
//"%b"     %o 0       %x  0xff




//json for go

type Student struct{
	ID int
	Name string
}


//结构体标签（定义结构体时） ----> Tag
//Tag 是结构体的元信息，可以在运行的时候通过反射的机制读取出来。
//格式： `key1:"value1" key2:"value2"`    key：表示包名
//json 表示在使用到json这个包的时候生效
type class struct{
	Title string  `json:"title" db:"title" xml:"title"`
	Students []Student `json:"student_list"`
}


func main(){
	cl := class{
		Title:"火箭",
		Students:make([]Student,0,20),   //可追加数据  cap容量20
	}
	for i := 0; i < 10; i++ {
		cl.Students = append(cl.Students,Student{
			ID:i,
			Name:fmt.Sprintf("%d",i),
		})
	}
	// 1. %v    只输出所有的值
	// 2. %+v 先输出字段名字，再输出该字段的值
	// 3. %#v 先输出结构体名字值，再输出结构体（字段名字+字段的值）
	// fmt.Printf("%#v",cl)

	//JSON序列化   Go语言中的数据 --> JSON字符串
	data,err :=json.Marshal(cl)
	if err !=nil{
		fmt.Println("失败")
	}
	fmt.Printf("%s\n",data)
	fmt.Printf("%T\n",data)

	//反序列化，params：一个字节数组，需要存放返回的值的引用地址
	var result class
	error := json.Unmarshal([]byte(data),&result)
	if error!=nil{
		fmt.Println("error...反序列化失败")
	}
	fmt.Printf("%#v",result)
}


