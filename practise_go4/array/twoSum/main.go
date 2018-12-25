package main

import "fmt"

func main() {
	var a = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(a)
	b := removeDuplicates(a)
	fmt.Println(b)
	fmt.Println(a)

}

func removeDuplicates(nums []int) int {
	var l = len(nums)
	if l <= 1 {
		return l
	}
	j := 1
	for i := j; i < l; i++ {
		if nums[i-1] != nums[i] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
