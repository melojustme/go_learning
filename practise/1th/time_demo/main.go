package main

import "time"
import "fmt"

func main() {
	timestamp := time.Now().Unix() //时间戳
	fmt.Println(timestamp)

	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05"))
}
