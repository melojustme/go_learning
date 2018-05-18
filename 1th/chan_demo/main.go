package main

import (
	"fmt"

	"sync"
)

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go Afun(i)
	}
	waitGroup.Wait()
}

var waitGroup sync.WaitGroup

func Afun(index int) {

	fmt.Println(index)

	waitGroup.Done()
}
