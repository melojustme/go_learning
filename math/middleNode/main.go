package main

import "fmt"

func main() {
	l.Print()
}

var l *LinkedList

func init() {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	l = &LinkedList{head: &ListNode{Next: n1}}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	var pre *ListNode = nil
	//

	return pre
}

type LinkedList struct {
	head *ListNode
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.Next
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
