package main

import "fmt"

func main() {

	sliceA := []int{1, 2, 3}
	printHelper("sliceA", sliceA)

	// 假设切片 slice 如下:
	slice := []int{1, 2, 3, 4, 5}
	newSlice := make([]int, len(slice))
	// 如何使用 copy 创建切片 newSlice, 该切片值为 [2, 3, 4]
	copy(newSlice, slice)
	printHelper("newSlice", newSlice)
}
func printHelper(name string, arr []int) {

	fmt.Println(name, "", arr)
}
