package main

import(
	"time"
	"fmt"
)

//time时间库的使用
func main() {
	now := time.Now()
	year := now.Year()
	//时间戳
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Println(now)
	fmt.Println(year)
	fmt.Println(timeStamp1)
	fmt.Println(timeStamp2)
}