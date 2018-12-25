package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := containsDuplicate(a)
	fmt.Println(b)
}

func containsDuplicate(nums []int) bool {
	var h = map[int]int{}
	for i := 0; i < len(nums); i++ {
		h[nums[i]] += 1
		if h[nums[i]] > 1 {
			return true
		}
	}
	return false
}
