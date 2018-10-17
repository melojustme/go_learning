package main

import (
	"fmt"
)

func main() {
	var nums = []int{11, 15, 2, 7}
	var target = 9
	// fmt.Println(nums, target)
	anwser := twoSum(nums, target)
	fmt.Println(anwser)
}

//
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i, v := range nums {
		sub := target - v
		fmt.Println(sub)
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}

	return nil
}