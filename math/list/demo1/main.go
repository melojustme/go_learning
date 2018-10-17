package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	e4 := l.PushBack(4)
	// fmt.Println(e4)
	// fmt.Println(e4.Value)
	e1 := l.PushFront(1)
	// fmt.Println(e1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
