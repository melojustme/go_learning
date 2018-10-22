package main

import (
	"fmt"
)

func main() {
	n6 := &ListNode{}
	n5 := &ListNode{Val: 5, Next: n6}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	l := &ListNode{Next: n1}
	l.Print()
	a := reverseList(l)
	a.Print()
	// l.Print()

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

func reverseList(head *ListNode) *ListNode {
	// cur := head
	// var root, last *ListNode
	// for cur != nil {
	// 	n := &ListNode{Val: cur.Val, Next: last}
	// 	if cur.Next == nil {
	// 		root = n
	// 	}
	// 	last = n
	// 	cur = cur.Next
	// }
	// return root

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	var pre *ListNode = nil
	cur := head.Next
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
