package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Add(-1)
			// defer wg.Done()

			EchoNumber(n)
		}(i)
		wg.Wait()
	}

}

func EchoNumber(i int) {
	time.Sleep(3e9)
	fmt.Println(i)
}
