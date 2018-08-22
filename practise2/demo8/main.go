package main

import "fmt"

func reverse(p *[]int) {
	for i, j := 0, len(*p)-1; i < len(*p)/2; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func main() {
	v := []int{1, 2, 3, 4}
	reverse(&v)
	fmt.Println(v)
}

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	person := Person{"vanyan", 21}
// 	// fmt.Printf("person<%s:%d>\n", person.name, person.age)
// 	person.sayHi()

// 	person.ModifyAge(28)

// 	person.sayHi()
// }

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) sayHi() {
// 	fmt.Printf("sayHi -- This is %s ,my age is %d\n", p.name, p.age)
// }

// func (p *Person) ModifyAge(age int) {
// 	fmt.Println("ModifyAge")
// 	p.age = age
// }
