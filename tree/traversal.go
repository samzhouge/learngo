package tree

import "fmt"

func (n *Node) Traverse() {
	n.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (n *Node) TraverseFunc(f func(*Node)) {
	if n == nil {
		return
	}
	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}
