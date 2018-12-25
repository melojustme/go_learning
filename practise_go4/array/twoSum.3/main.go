package main

import "fmt"

func main() {
	a := []int{-1, -100, 3, 99}
	k := 2
	rotate(a, k)
	fmt.Println(a)
}

func rotate(nums []int, k int) {
	var l = len(nums)
	for k > 0 {
		for i := l - 1; i > 0; i-- {
			// fmt.Println(i, nums[i])
			nums[i-1], nums[i] = nums[i], nums[i-1]
		}
		k--
	}
	// fmt.Println("111", nums)
}
