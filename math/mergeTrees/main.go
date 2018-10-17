package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var t *TreeNode
	if t1 == nil || t2 == nil {
		return nil
	}

	return t
}
