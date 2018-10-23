package main

import (
	"fmt"
)

func main() {

	var nums1 = []int{4, 9, 5}
	var nums2 = []int{9, 4, 9, 8, 4}
	a := intersection2(nums1, nums2)
	fmt.Println(a)
}
func intersection(nums1 []int, nums2 []int) []int {
	a := []int{}
	var k = make(map[int]int)
	for _, v := range nums1 {
		for _, b := range nums2 {
			if b == v {
				// fmt.Println(b)
				k[v] = v

			}
		}
	}
	// fmt.Println(k)
	for _, v := range k {
		a = append(a, v)
	}
	return a

}
func intersection2(nums1 []int, nums2 []int) []int {
	var a = []int{}
	var b = make(map[int]int)
	for _, v := range nums1 {
		b[v] = 1
	}
	for _, v := range nums2 {
		if b[v] == 1 {
			a = append(a, v)
			b[v] = 2
		}

	}
	return a
}
