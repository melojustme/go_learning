package main

import (
	"fmt"
)

func main() {
	var a = 17
	fmt.Println(a, addDigits(a))
}

func addDigits(num int) int {
	return (num-1)%9 + 1
}
