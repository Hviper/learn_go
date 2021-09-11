package main


import (
	"fmt"
	"time"
	"sync"
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



//Try to encapsulate a custom array method
type myList []string
func (list *myList)add(val string){
	*list = append(*list,val)
}

// list := myList{
// 	"name",
// 	"age",
// }
// list.add("school")
// fmt.Println(list)

func traverse(person []string ){
	for _, name := range person {
		fmt.Println("name :", name)
	}
}


//defer ----> finall(最终执行的语句)   比如文件的close事件需要无论如何都执行
// defer fmt.Println("end....")，
//多个defer的 话是根据压栈的形式来进行执行




//高并发
// var wg sync.WaitGroup
// wg.Add(2)
// go func(){
// 	count(5,"🐏")
// 	wg.Done()
// }()
// go func(){
// 	count(3,"🐂")
// 	wg.Done()
// }()
// wg.Wait()
func current(){
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		count(5,"🐏")
		wg.Done()
	}()
	go func(){
		count(3,"🐂")
		wg.Done()
	}()
	wg.Wait()
}
func count(n int,animal string){
	for i := 0; i < n; i++ {
		fmt.Println(i+1,animal)
		time.Sleep(time.Second*3)
	}
}


//change struct only by pointer
// var mystruct MyStruct = MyStruct{
// 	name:"apache",
// 	age:19,
// 	school:"清华",
// }
// changeStruct(&mystruct)
// fmt.Println(mystruct)
type MyStruct struct{
	name string 
	age int 
	school string
}
func changeStruct(s *MyStruct){
	s.name = "kiwi"
	s.age = 18

}

	//new关键字实现创建返回一个结构体的指针，指向当前创建出来的对象
	// m := new(MyStruct)
	//结构体中可以通过地址直接访问属性（语法糖）
	// (*m).name = "kiwi"
	// changeStruct(&mystruct)
	// fmt.Println(mystruct)
	// fmt.Println(*m)


//闭包，返回一个func函数
//a()()
func a() func(){
	return func(){
		fmt.Println("内部执行成功")
	}
}


//构造函数, ==> 工厂模式(内部实例化一个对象，一般返回指针类型，也可返回对象)
//当结构体比较复杂的时候（内存字段比较多），建议返回一个地址（地址长度固定），而不是返回一个值拷贝来的对象，从而降低性能开销
//--------------------------------->
//返回为 一个地址
// s := newMyStruct("kiwi",18,"北京大学")
// fmt.Println(s)
func newMyStruct(name string,age int,school string) *MyStruct{
	//返回实例对象的地址 （& 取址符）
	return &MyStruct{
		name:name,
		age:age,
		school:school,
	}
}


//Go语言基础14 方法和接收者
type Person struct{
	name string
	age int
}
//为Person创建构造函数  (*Person：指针类型)
func newPerson(name string,age int) *Person{
	//取址
	return &Person{
		name:name,
		age:age,
	}	
}

//为Person类型定义的方法   【函数和方法的区别，函数不是隶属于特定的对象，而是一个独立的存在，方法是争对某个特定的对象而言的】
//下面两行效果一致
// (*point).run()
// point.run()
func (p Person) run(){
	fmt.Println("--->",p.name)
}


//修改年龄的方法,需要接受一个指针类型 ，这样才能修改生效
//指针接收者，指的是接收者的类型是指针

func (p *Person) setAge(age int){
	p.age = age
}

//这个方法调用无法修改结构体内部的值，因为结构体默认是值传递，【值拷贝 ！！，带来的另一个问题就是内存开销问题】
//什么时候应该使用指针类型？
//1.需要修改接收者的值
//2.接收者是拷贝代价比较大的对象
//3.保证一致性，如果有某个方法使用指针接收者，那么其他的方法也应该使用指针接收者
func (p Person) setAge2(age int){
	p.age = age
}



//结构体的匿名字段名称，外部通过可以类型来查找(局限性很大，只能通过类型判断)
// a := TestClass{
// 	"kiwi",
// 	18,
// }
// fmt.Println(a)
//访问：【通过类型来查找，下面的嵌套中例子是通过结构体类型来查找a.Address】
//fmt.Println(a.string)
//fmt.Println(a.int)
type TestClass struct{
	string
	int
}


//嵌套（可以使用匿名嵌套字段结构体）
type Address struct{
	Province string
	City string
}
type School struct{
	name string
	old int
	//匿名字段
	Address
}

func main(){
	address := Address{
		Province:"中国省会",
		City:"北京",
	}
	school := School{
		name:"北京大学",
		old :100,
		Address:address,
	}
	fmt.Println(school)
	//访问(两个方法等效)
	fmt.Println(school.Address.Province)
	fmt.Println(school.Province)

}