package main

import "fmt"

func main() {
	a := []int{1, 2, 2, 1}
	b := []int{2, 2}

	c := intersect(a, b)
	fmt.Println(c)
}

func intersect(n1 []int, n2 []int) []int {
	var c = []int{}
	var h1, h2 = map[int]int{}, map[int]int{}
	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			if n1[i] == n2[j] && h1[i] != 1 && h2[j] != 1 {
				c = append(c, n1[i])
				h1[i] = 1
				h2[j] = 1
			}
		}
	}
	return c
}
