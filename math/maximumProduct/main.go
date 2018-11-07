package main

import (
	"fmt"
	"sort"
)

func main() {

	var i int = -20
	fmt.Println(^i, ^i+1)

	var c = []int{999, 999, 999}
	fmt.Println(c)
	a := maximumProduct(c)
	fmt.Println("答案", a)
}

func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	var h = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 { //如果是负数补码
			h[^nums[i]+1] = nums[i]
		} else {
			h[nums[i]] = nums[i]
		}
	}
	// fmt.Println(h)
	var n2 []int
	for k, _ := range h {
		n2 = append(n2, k)
	}
	// fmt.Println(n2)
	j := sort.IntSlice(n2)
	if !sort.IsSorted(j) {
		sort.Sort(j)
	}
	// fmt.Println(j)
	// a := nums
	b := len(j)
	fmt.Println(h[j[b-1]])
	fmt.Println(h[j[b-1]] * h[j[b-2]])
	return h[j[b-1]] * h[j[b-2]] * h[j[b-3]]
	// return a[b-1] * a[b-2] * a[b-3]
}
