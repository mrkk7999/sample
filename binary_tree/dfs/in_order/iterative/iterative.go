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

type stack struct {
	elements []interface{}
}

func createEmptyStack() interface{} {
	return &stack{
		elements: nil,
	}
}
func (s *stack) push(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *stack) pop() {
	s.elements = s.elements[:len(s.elements)-1]
}
func (s *stack) isEmpty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}
func (s *stack) top() interface{} {
	return s.elements[0]
}
func (root *TreeNode) inOrder() []int {
	var res []int
	var stk stack
	//stk.push(root)
	var currentNode = root
	//res = append(res, currentNode.NodeVal)
	//for stk.isEmpty() {
	//	if currentNode.LeftNode != nil {
	//		stk.push(currentNode.LeftNode)
	//	}
	//	currentNode = stk.top().(*TreeNode)
	//	res = append(res, currentNode.NodeVal)
	//	if currentNode.RightNode != nil {
	//		stk.push(currentNode.RightNode)
	//	}
	//}
	var flag = 0
	for flag == 0 {
		if currentNode != nil {
			stk.push(currentNode)
			currentNode = currentNode.LeftNode
		} else {
			if stk.isEmpty() {
				flag = 1
			}
			currentNode = stk.top().(*TreeNode)
			stk.pop()
			res = append(res, currentNode.NodeVal)
			currentNode = currentNode.RightNode
		}
	}
	return res
}
func main() {
	root := InitTree()
	fmt.Println(root.inOrder())
}
