package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	a := make([]int, 10)
	for k, _ := range a {
		rand.Seed(time.Now().UnixNano())
		a[k] = rand.Intn(100)
		// fmt.Println(a[k])
	}
	fmt.Println(a)
	// for i := 0; i < len(a); i++ {
	// 	rand.Seed(time.Now().UnixNano())
	// 	a[i] = rand.Int()
	// 	fmt.Println(a[i])
	// }
	// fmt.Println(a)
	//
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] > a[j] {

			}
		}
	}
	fmt.Println(a)

}
