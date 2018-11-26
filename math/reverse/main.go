package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := 1112222
	b := reverse(a)
	fmt.Println(b)
	// s := "1234567890"
	// fmt.Println(reverseString(s))
}
func reverse(x int) int {

	//
	if x < 0 {
		y := ^x + 1 //
		a := 0
		for y != 0 {
			rem := y % 10
			a = a*10 + rem
			y = y / 10
		}
		if a*-1 > 2147483647 || a*-1 < -2147483648 {
			return 0
		}
		return a * -1
	}
	//
	if x > 0 {
		y := x
		a := 0
		for y != 0 {
			rem := y % 10
			a = a*10 + rem
			y = y / 10
		}
		if a > 2147483647 || a < -2147483648 {
			return 0
		}
		return a
	}

	return 0
}

func reverse2(x int) int {

	if x > 0 {
		a := strconv.Itoa(x)
		b := reverseString(a)
		c, _ := strconv.Atoi(b)
		if c > 2147483647 || c < -2147483648 {
			return 0
		}
		return c
	} else if x < 0 {
		a := strconv.Itoa(x)
		b := reverseString(a[1:])
		// fmt.Println(a, b)
		c, _ := strconv.Atoi("-" + b)
		// fmt.Println(c)
		if c > 2147483647 || c < -2147483648 {
			return 0
		}
		return c
	}
	return 0
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
