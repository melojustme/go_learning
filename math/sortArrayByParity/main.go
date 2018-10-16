package main

import "fmt"

func main() {
	var a = []int{4, 2, 5, 7}
	fmt.Println(a)
	b := sortArrayByParityII(a)
	fmt.Println(b)

}
func sortArrayByParityII(A []int) []int {
	var a = make([]int, 0)
	var b = make([]int, 0)
	for _, v := range A {
		if v%2 != 0 {
			b = append(b, v)
		} else {
			a = append(a, v)
		}
	}
	a = append(a, b...)
	return a
}
