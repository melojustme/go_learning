package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums1 = []int{1, 3}
	var nums2 = []int{2, 4}
	a := findMedianSortedArrays(nums1, nums2)
	fmt.Println(a)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for i := 0; i < len(nums2); i++ {
		nums1 = append(nums1, nums2[i])
	}
	sort.Ints(nums1)
	l := len(nums1)
	if l%2 == 0 {
		return (float64(nums1[l/2-1]) + float64(nums1[l/2])) / 2
	} else {
		fmt.Println("a")
		return float64(nums1[l/2])
	}
}
