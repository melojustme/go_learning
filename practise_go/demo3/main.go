package main

import (
	"fmt"
)

func main() {
	var a = []int{3000, 5000, 10000, 2000, 200000}
	a1, a2, a3 := printPrime(a)
	fmt.Println(a1, a2, a3)
}

//数据分组
func printPrime(a []int) (d, b, c []int) {
	// var b, c, d []int
	var a1 = make([]int, 0)
	var a2 = make([]int, 0)
	var a3 = make([]int, 0)

	for _, v := range a {
		if v >= 3000 && v < 5000 {
			a1 = append(a1, v)
		}
		if v >= 5000 && v < 10000 {
			a2 = append(a2, v)
		}
		if v >= 10000 {
			a3 = append(a3, v)
		}
	}
	return a1, a2, a3
}
