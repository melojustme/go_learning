package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Circle) area() float64 {
	return r.radius * r.radius * math.Pi
}

/**
虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
method里面可以访问接收者的字段
调用method通过.访问，就像struct里面访问字段一样

*/
func main() {
	r1 := Rectangle{12, 2}

	r2 := Circle{9}

	fmt.Println(r1.area()) //矩形的面积
	fmt.Println(r2.area()) //圆的面积

}
