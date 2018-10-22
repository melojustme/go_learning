package main

func main() {

}

type Node struct {
	data Object
	next *Node
}
type Object interface{}

type List struct {
	size uint64
	head *Node
	tail *Node
}

func (list *List) Init() {
	(*list).size = 0
	(*list).head = nil
	(*list).tail = nil
}

func (list *List) Append(node *Node) {
	(*list).size = 1
	(*list).head = node
	(*list).tail = node
}
