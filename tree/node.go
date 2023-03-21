package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(v int) *Node {
	return &Node{Value: v}
}

func (n Node) Print() {
	fmt.Printf("%d ", n.Value)
}

func (n *Node) SetValue(v int) {
	if n == nil {
		fmt.Println("set nil, ignore")
	}
	n.Value = v
}
