package main
import (
	"fmt"
)




func add(a map[string]string){
	a["name"] = "ppp"
	fmt.Println("test")
}

func main(){
	//支持引用
	x := make(map[string]string)
	add(x)
	fmt.Println(x)

	var n MyFloat = 1.5
	n.show()

}
