package main

import (
	"fmt"
)

func main() {
	var nums = []int{4, 1, 2, 1, 2}
	a := singleNumber(nums)
	fmt.Println(a)
}
func singleNumber2(nums []int) int {

	var a int
	var b = make(map[int]int)
	for _, v := range nums {
		for _, s := range nums {
			if v == s {
				if b[v] == 1 {
					b[v] = 2
				} else {
					b[v] = 1
				}
			}
		}
	}

	for k, v := range b {
		if v == 1 {
			a = k
		}
	}
	return a
}
func singleNumber(nums []int) int {
	res := 0

	for _, element := range nums {
		res ^= element
	}

	return res
}
