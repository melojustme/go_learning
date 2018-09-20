package main

import (
	"fmt"
)

//递归
func main() {
	var i int = 5

	fmt.Printf("%d\n", Factorial(uint64(i)))
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
