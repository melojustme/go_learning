package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// value := counter
		// atomic.AddInt64(&counter, 1)
		// runtime.Gosched()
		// value++
		// counter = value
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
