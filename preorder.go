package main


type Node struct {
     Val int
     Children []*Node
}

var res []int

func preorder(root *Node) []int {
	res = make([]int, 0)
	pre(root)

	return res
}

func pre(node *Node) {
	if node == nil {
		return
	}
	res = append(res, node.Val)
	for _,v := range node.Children {
		pre(v)
	}
	return
}