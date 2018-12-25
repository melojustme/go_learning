package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := maxProfit(a)
	fmt.Println("ZHI ", b)
	fmt.Println(a)

}

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	var j = 1
	var a int
	for i := j; i < l; i++ {
		if prices[i] > prices[i-1] {
			a += prices[i] - prices[i-1]
			// fmt.Println("a", a)
		}
	}
	return a
}
