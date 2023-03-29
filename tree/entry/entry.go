package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	*tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")
}

func main() {
	//root := tree.Node{Value: 3}
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = tree.CreateNode(0)
	root.Left = new(tree.Node)
	root.Left.Right = &tree.Node{Value: 2}

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	//var pRoot *tree.Node
	//fmt.Println(pRoot)

	fmt.Printf("%T\n", root)

	fmt.Print("root traversal: ")
	root.Traverse()

	fmt.Print("In-order traversal: ")
	root.Node.Traverse()
	fmt.Println()

	fmt.Print("My own post-order traversal: ")
	root.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
