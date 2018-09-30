package main

import (
	"fmt"
)

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7}
	// for i := range n {
	// 	if i == 3 {
	// 		fmt.Println("a:", i)
	// 		n[i] |= i
	// 	}
	// }
	for i := 0; i < len(n); i++ {
		if i == 6 {
			fmt.Println("a:", i)
			n[i] |= i

		}
		fmt.Println("n[i]:", n[i])

	}

	fmt.Println("b:", n)

}
