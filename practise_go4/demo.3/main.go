package main

import "fmt"

func main() {
	//O(1)
	i := 8
	j := 9
	sum := i + j
	fmt.Println(sum)

	fmt.Println(cal(3))

	fmt.Println(logn(12))
}

//O(n)
func cal(n int) int {
	sum := 0
	i := 1
	for ; i <= n; i++ {
		sum += i
	}
	return sum
}

//O(logn)
func logn(n int) int {
	i := 1
	for i <= n {
		i = i * 2
		fmt.Println("hdhdhhd ", i)
	}
	return i
}

//O(m+n)
func cal2(m, n int) int {

	sum_1 := 1
	i := 1
	for ; i < m; i++ {
		sum_1 = sum_1 + i
	}

	sum_2 := 1
	j := 1
	for ; j < m; j++ {
		sum_2 = sum_2 + j
	}
	return sum_1 + sum_2
}
