package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := -1234
	b := reverse(a)
	fmt.Println(b)

}

func reverse(x int) int {
	if x > 0 {
		b := strconv.Itoa(x)
		x, _ = strconv.Atoi(reverseString(b))
	}
	if x < 0 {
		b := strconv.Itoa(x)
		x, _ = strconv.Atoi("-" + reverseString(b[1:]))
	}
	if x < -2147483648 || x >= 2147483648 || x == 0 {
		return 0
	}
	return x
}

func reverseString(s string) string {
	var a []byte
	for i := 0; i < len(s); i++ {
		a = append(a, s[i])
	}
	var b []byte
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	return string(b)
}
