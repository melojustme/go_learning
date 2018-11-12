package main

import (
	"fmt"
	"strconv"
)

func main() {
	// a := 38
	a := 23465
	b := addDigits(a)
	fmt.Println("答案", b)

}

func addDigits(a int) int {
	l := len(strconv.Itoa(a))
	// k := []int{}
	c := 0
	for l > 0 {
		// k = append(k, a/pow(10, l-1)%10)
		c += a / pow(10, l-1) % 10
		l--
	}
	// fmt.Println(k)
	// fmt.Println(c)
	if c < 10 {
		return c
	}
	return addDigits(c)
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
