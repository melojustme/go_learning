package main

import (
	"fmt"
)

func main() {

	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	fmt.Println(arr1)

	//2
	arr1 = [5]int{1, 2, 3, 4, 5}
	printHelper("tmp", arr1)

	//3

}

func printHelper(name string, arr [5]int) {

	fmt.Println(name, "", arr)
}
