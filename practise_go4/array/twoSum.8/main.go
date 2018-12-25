package main

import "fmt"

func main() {
	a := []int{0, 1, 0, 2, 3, 8}

	moveZeroes(a)
	fmt.Println(a)
}

func moveZeroes(nums []int) {
	var l = len(nums)
	var a int
	for i := 1; i < l; i++ {
		if nums[i-1] == 0 {
			a += 1
		}
	}
	// fmt.Println(a)
	for i := 0; i < a; i++ {
		for i := 1; i < l; i++ {
			if nums[i-1] == 0 {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		}
	}
}
