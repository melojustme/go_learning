package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		// go func() {
		// fmt.Println(i)
		go incr()
		// }()

	}
	time.Sleep(time.Second * 1)

}

var count = 0
var lock sync.Mutex

func incr() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Println(count)
}
