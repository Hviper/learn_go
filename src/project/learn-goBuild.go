package project
import(
	"fmt"
)

func test(){
	
	ad := struct{
		name string
		age int 
	}{
		name:"kiwi",
		age:18,
	}
	fmt.Println(ad)
}
