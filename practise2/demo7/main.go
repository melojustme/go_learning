package main

import "fmt"

func main() {
	// array := [...]string{2: "1", 3: "1"}
	// fmt.Println(array)
	// fmt.Println(len(array))
	// fmt.Println(cap(array))
	slice := make([]string, 0)
	// slice[1] = "1"
	fmt.Println(slice)
	slice = append(slice, "333")
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
