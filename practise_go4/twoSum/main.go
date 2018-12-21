package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 13
	a := twoSum(nums, target)
	fmt.Println(a)
}

func twoSum(nums []int, target int) []int {
	var a []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				a = append(a, i)
				a = append(a, j)
				return a
			}
		}
	}
	return nil
}
