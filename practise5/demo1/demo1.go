package main

import (
	"fmt"
	"time"
)

//map
func main() {
	// countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	// fmt.Println(countryCapitalMap)

	// for a := range countryCapitalMap {
	// 	// fmt.Println(a, v)
	// 	fmt.Println(countryCapitalMap[a])
	// }
	// delete(countryCapitalMap, "France")
	// fmt.Println(countryCapitalMap)
	// var badMap2 = map[interface{}]int{
	// "1": 1,
	// []int{2}: 2,
	// 3:        3,
	// }
	// fmt.Println(badMap2)

	// var a = map[bool]string{true: "123", false: "33333"}
	// a[true] = "112"
	// a[false] = "333"
	// fmt.Println(a)

	// var m1 map[string]string
	// m1 = make(map[string]string)
	// m1["1111"] = "22222"
	// fmt.Println(m1)

	// m := map[int]string{1: "1"}
	// go read(m)
	// time.Sleep(time.Second)

	// go write(m)
	// time.Sleep(30 * time.Second)

}
func read(n map[int]string) {
	for {
		fmt.Println(n)
		time.Sleep(1)
	}
}
func write(n map[int]string) {
	for {
		fmt.Println(n)
		time.Sleep(1)
		n[1] = "write"
	}
}
