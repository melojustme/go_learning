package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1
	var sum int
	n := 100
	//bad way
	t1 := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		sum += a + i
	}
	fmt.Println(sum)
	t2 := time.Now().UnixNano()
	fmt.Println("时间", t2-t1)
	//good  way
	t3 := time.Now().UnixNano()
	sum = (1 + n) * n / 2
	fmt.Println(sum)
	t4 := time.Now().UnixNano()

	fmt.Println("时间", t4-t3)

}
