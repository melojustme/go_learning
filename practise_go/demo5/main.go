package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	wg.Add(2)
	go dowork("A")
	go dowork("B")

	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

var (
	shutdown int64
	wg       sync.WaitGroup
)

func dowork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Dowm\n", name)
			break
		}
	}

}
