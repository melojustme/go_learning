package main

import (
	"fmt"
	"sync"
)

func main() {

	for i := 0; i < 10; i++ {

		waitgroup.Add(1)
		go func(n int) {
			// defer waitgroup.Done()
			test(n)
		}(i)
		waitgroup.Wait()

	}
	fmt.Println("Done!")

}

var waitgroup sync.WaitGroup

func test(shownum int) {
	fmt.Println(shownum)

}
