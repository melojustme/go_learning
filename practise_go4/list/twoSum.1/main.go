package main

import "fmt"

func main() {
	printLink(NewNodes())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var l *ListNode

	return l
}

func NewNodes() *ListNode {
	head := &ListNode{0, nil}
	for i := 10; i > 0; i-- {
		// fmt.Println(i)
		a := &ListNode{i, nil}
		a.Next = head
		head.Next = a
	}
	return head
}

func printLink(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
