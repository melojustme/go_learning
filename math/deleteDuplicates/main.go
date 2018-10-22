package main

import (
	"fmt"
)

func main() {
	n5 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 2, Next: n5}
	n3 := &ListNode{Val: 2, Next: n4}
	n2 := &ListNode{Val: 1, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	// l := &ListNode{Next: n1}
	// fmt.Println(n1, n2, n3, n4, n5)
	l := &ListNode{Next: n1}
	l.Print()
	//
	deleteDuplicates(l)
	// fmt.Println(a)
	// a.Print()
	l.Print()
}

//打印链表
func (this *ListNode) Print() {
	cur := this.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.Val)
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var a *ListNode
	a = head
	for a != nil && a.Next != nil {
		fmt.Println(a.Val)
		fmt.Println(a.Next.Val)

		if a.Val == a.Next.Val {
			a.Next = a.Next.Next
		} else {
			a = a.Next
		}
	}

	return head
}
