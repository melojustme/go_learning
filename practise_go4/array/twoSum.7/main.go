package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 8}

	c := plusOne(a)
	fmt.Println(c)
}

func plusOne(digits []int) []int {
	var c []int
	l := len(digits)
	var b int
	// for i := l - 1; i > 1; i-- {
	// 	fmt.Println(i)
	// 	fmt.Println(pow(10, i))

	// }
	for i := 0; i < l; i++ {
		// digits[i] * pow(10, l)
		// l--
	}
	fmt.Println("b", b)
	return c
}

func pow(x, n int) int {
	ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
