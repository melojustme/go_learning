package main

import (
	"container/list"
	"fmt"
)

func main() {
	w := list.New()
	w.PushBack("1")
	w.PushBack(1)
	w.PushBack(2)
	w.PushBack("4df")
	for e := w.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(w)
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	return head
}
