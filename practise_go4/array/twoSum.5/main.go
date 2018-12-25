package main

import "fmt"

func main() {
	a := []int{1, 1, 2, 2, 3, 3, 4}
	b := singleNumber(a)
	fmt.Println(b)
}

func singleNumber(nums []int) int {
	var h = map[int]int{}
	for i := 0; i < len(nums); i++ {
		h[nums[i]] += 1
	}
	for k, v := range h {
		if v == 1 {
			return k
		}
	}
	return 0
}
