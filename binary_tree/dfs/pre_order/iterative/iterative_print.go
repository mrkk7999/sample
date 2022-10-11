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

type Stack struct {
	elements []interface{}
}

func (s *Stack) push(element interface{}) {
	s.elements = append(s.elements, element)
}
func (s *Stack) pop() {
	s.elements = s.elements[:len(s.elements)-1]
}
func (s *Stack) isEmpty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}
func (s *Stack) size() int {
	return len(s.elements)
}
func (s *Stack) top() interface{} {
	return s.elements[len(s.elements)-1]
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
func (root *TreeNode) preOrder() []int {
	var res []int
	var stack Stack
	stack.push(root)
	for !stack.isEmpty() {
		currentNode := stack.top().(*TreeNode)
		res = append(res, currentNode.NodeVal)
		stack.pop()
		if currentNode.RightNode != nil {
			stack.push(currentNode.RightNode)
		}
		if currentNode.LeftNode != nil {
			stack.push(currentNode.LeftNode)
		}
	}
	return res
}
func main() {
	root := InitTree()
	fmt.Println(root.preOrder())
}
