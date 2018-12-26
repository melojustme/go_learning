package main

import "fmt"

func main() {
	var a *ListNode
	a.Val = 1
	a.Next.Val = 2
	a.Next.Next.Val = 3
	a.Next.Next.Next.Val = 3
	fmt.Println(a)
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	return head
}
