package node

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) AddListNode(val int) {
	for this.Next != nil { //如果下一个不是空，赋值成下一个
		this = this.Next
	}
	//this.Next 是nil的时候，就是最后一个节点
	fmt.Println("this.Next", this.Next)
	this.InsetAfter(this, val) //赋值给当前值
}

func (this *ListNode) InsetAfter(p *ListNode, val int) {
	if p == nil {
		return
	}
	newNode := NewNode(val) //声明一个ListNode
	oldNext := p.Next       //oldNext获取p的next
	p.Next = newNode        //
	newNode.Next = oldNext
	return
}

func NewNode(l int) *ListNode {
	return &ListNode{l, nil}
}

func (this *ListNode) PrintLink() {
	format := ""
	cur := this.Next
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.Val)
		cur = cur.Next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)

}
