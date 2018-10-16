package main

import "fmt"

func main() {
	var a = []int{3, 1, 2, 4}
	fmt.Println(a)
	b := sortArrayByParity(a)
	fmt.Println("结果", b)

}
func sortArrayByParity(A []int) []int {
	var a = make([]int, 0)
	var b = make([]int, 0)
	var c = make([]int, len(A))

	for _, v := range A {
		if v%2 != 0 {
			b = append(b, v)
		} else {
			a = append(a, v)
		}
	}
	var i1, i2 int
	for k, _ := range c {
		if k%2 != 0 {
			c[k] = b[i1]
			i1++
		} else {
			c[k] = a[i2]
			i2++
		}
	}
	return c
}
