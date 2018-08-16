package main

import "fmt"

type person struct {
	name   string
	age    byte
	isDead bool
}

func main() {
	p1 := &person{name: "zzy", age: 100}
	p2 := &person{name: "dj", age: 99}
	p3 := &person{name: "px", age: 20}
	people := []*person{p1, p2, p3}
	whoIsDead(people)
	// fmt.Println(people)
	for _, p := range people {
		// fmt.Println(p)
		if p.isDead {
			fmt.Println("who is dead?", p.name)
		}
	}
}

func whoIsDead(people []*person) {
	// fmt.Println(people)
	for _, p := range people {
		if p.age < 50 {
			p.isDead = true
		}
		fmt.Println(p)
	}
	fmt.Println(people)

}
