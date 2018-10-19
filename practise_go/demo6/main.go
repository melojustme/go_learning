package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	done := make(chan struct{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, done)
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doIt(workerID int, done <-chan struct{}) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerID)
}
