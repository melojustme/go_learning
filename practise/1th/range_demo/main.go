package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)

	//
	// v := []int{1, 2, 3}
	// for i := 0; i < len(v); i++ {
	// 	v = append(v, i)
	// 	if len(v) > 100 {
	// 		break
	// 	}
	// }
	// fmt.Println(v)

}
