package main

import (
	"fmt"
)

func main() {
	// nums1 := []int{1, 2}
	// nums2 := []int{1, 1}
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println("kkkk", nums1, nums2)
	a := intersect(nums1, nums2)
	fmt.Println("da", a)

}
func intersect(nums1 []int, nums2 []int) []int {
	var num []int
	var h = make(map[int]int)
	for _, v := range nums1 {
		h[v] += 1
	}
	fmt.Println(h)
	for _, a := range nums2 {
		if h[a] > 0 {
			num = append(num, a)
			h[a] -= 1
		}
	}
	return num
}
