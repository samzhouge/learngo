package tree

func (n *Node) Traverse() {
	if n == nil {
		return
	}
	n.Left.Traverse()
	n.Print()
	n.Right.Traverse()
}
