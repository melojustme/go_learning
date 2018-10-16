package main

import "fmt"

func main() {
	x, y := 1, 4
	b := hammingDistance(x, y)
	fmt.Println("结果", b)

}
func hammingDistance(x int, y int) int {
	var d int

	for {
		if x == 0 && y == 0 {
			break
		}
		fmt.Println("x", x&1, "y", y&1)

		if x&1 != y&1 {
			d++
			fmt.Println(d)
		}
		x = x >> 1
		y = y >> 1
		fmt.Println("x>>", x, "y>>", y)

	}
	return d
}
