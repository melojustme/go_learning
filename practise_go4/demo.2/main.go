package main

import "fmt"

func main() {
	var a [5]int
	var b = [5]int{1, 2, 3, 45, 5}
	a = b
	fmt.Println(a)
}
