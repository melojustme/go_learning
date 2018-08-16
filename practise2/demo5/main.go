package main

import (
	"fmt"
	"math/rand"
	"time"
)

//1.
// func main() {
// 	http.Handle("/", http.)
// 	http.ListenAndServe(":8080", nil)
// }
//2.
func producer(header string, channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31()) // 将随机数和字符串格式化为字符串发送给通道
		time.Sleep(time.Second)
	}
}

func customer(channel <-chan string) {
	for {
		message := <-channel // 从通道中取出数据, 此处会阻塞直到信道中返回数据
		fmt.Println(message)
	}
}

func main() {
	channel := make(chan string) //创建一个字符串类型的通道
	go producer("cat", channel)
	go producer("dog", channel)
	customer(channel)
}
