package main

import "fmt"

type TreeNode struct {
	NodeVal   int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{
		NodeVal:   value,
		LeftNode:  nil,
		RightNode: nil,
	}
}

func InitTree() *TreeNode {
	root := CreateNode(1)
	// left subtree
	root.LeftNode = CreateNode(2)
	root.LeftNode.LeftNode = CreateNode(4)
	root.LeftNode.RightNode = CreateNode(5)
	// right subtree
	root.RightNode = CreateNode(3)
	root.RightNode.LeftNode = CreateNode(6)
	root.RightNode.RightNode = CreateNode(7)
	return root
}
func (root *TreeNode) preOrder() {
	if root == nil {
		return
	}
	fmt.Print(root.NodeVal, " ")
	root.LeftNode.preOrder()
	root.RightNode.preOrder()
}

func main() {
	root := InitTree()
	root.preOrder()
}
