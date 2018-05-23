package main

import (
	"fmt"
	"strings"
)

func main() {
	arr1 := []string{"2", "3", "333"}
	arr2 := []string{"2", "3", "111"}
	arr3 := Compare(arr1, arr2)
	fmt.Println(arr3)

}

func Compare(arr1 []string, arr2 []string) (arr3 []string) {

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			bl := strings.EqualFold(arr1[i], arr2[i])
			if !bl {
				fmt.Println(arr1[i], arr2[i])
				arr3 = append(arr3, arr1[i])
				arr3 = append(arr3, arr2[i])
				break
			}
		}
	}
	return
}
