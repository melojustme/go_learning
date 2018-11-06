package main

import (
	"fmt"
	"math"
)

func main() {
	x := 8
	i := mySqrtmySqrt(x)
	fmt.Println(i)
}
func mySqrtmySqrt(x int) int {
	a := math.Sqrt(float64(x))
	return int(a)
}
