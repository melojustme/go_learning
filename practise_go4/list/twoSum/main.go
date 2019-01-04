package main

import (
	"fmt"
	"go_learning/practise_go4/list/twoSum/node"
)

func main() {
	w := &node.ListNode{0, nil}
	w.AddListNode(1)
	w.AddListNode(4)
	w.AddListNode(3)
	fmt.Println(w)
	w.PrintLink()

}
