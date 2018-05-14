package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y = 3, 4
	var f float64 = math.Sqrt(float64(x + y))
	fmt.Println(f)
	fmt.Println(float64(x + y))
}
