package main

import (
	"fmt"
	"time"
)

func main() {

	var a = []int{1, 2, 3, 4, 5, 6, 34, 2, 2, 3, 4, 4, 23, 2, 2, 2, 2, 2, 1, 3, 4, 1, 4, 4, 2, 4, 1}
	var ch = make(chan bool, len(a))
	t1 := time.Now().UnixNano()
	fmt.Println(t1)
	for i := 0; i < len(a); i++ {
		go func(k int) {
			fmt.Println(a[k])
		}(i)
		ch <- true
	}
	for i := 0; i < len(a); i++ {
		<-ch
		time.Sleep(time.Millisecond * 50)

	}
	t2 := time.Now().UnixNano()
	fmt.Println(t2 - t1)

}
