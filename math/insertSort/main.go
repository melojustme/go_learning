package main

import (
	"fmt"
	"sort"
)

func main() {
	// a := []int{4, 5, 6, 3, 2, 1}
	// InsertSort2(a)
	// fmt.Println(a)
	a := 15
	GuessingGame(a)
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

func InsertSort2(a []int) {
	sort.Ints(a)
}
func SortSearch(a []int) {
	// sort.Search()
}
func GuessingGame(a int) {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(a, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
