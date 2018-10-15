package main

import (
	"fmt"
)

func main() {
	var nums = []int{11, 15, 2, 7}
	var target = 9
	// fmt.Println(nums, target)
	anwser := TwoSum2(nums, target)
	fmt.Println(anwser)
}

func TwoSum2(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
		fmt.Println(m)
	}

	return nil
}
