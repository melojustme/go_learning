package main

import (
	"fmt"
)

func main() {
	// defer_call()
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		fmt.Println(stu)
		m[stu.Name] = &stu
		fmt.Println(m[stu.Name])

	}
	fmt.Println(m)
}

// func defer_call() {
// 	defer func() { fmt.Println("打印前") }()
// 	time.Sleep(1 * time.Second)
// 	defer func() { fmt.Println("打印中") }()
// 	time.Sleep(1 * time.Second)
// 	defer func() { fmt.Println("打印后") }()
// 	time.Sleep(1 * time.Second)
// 	panic("触发异常")
// }
