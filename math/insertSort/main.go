package main

import (
	"fmt"
)

func main() {
	a := []int{4, 5, 6, 3, 2, 1}
	InsertSort(a)
	fmt.Println(a)
}

func InsertSort(a []int) {
	arrLen := len(a)
	if arrLen <= 1 {
		return
	}
	for i := 1; i < arrLen; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = v
	}
}
