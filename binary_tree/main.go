package main

import (
	"fmt"
)

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

//func InitTree() *TreeNode {
//	root := CreateNode(1)
//	// left subtree
//	root.LeftNode = CreateNode(2)
//	root.LeftNode.LeftNode = CreateNode(4)
//	root.LeftNode.RightNode = CreateNode(5)
//	// right subtree
//	root.RightNode = CreateNode(3)
//	root.RightNode.LeftNode = CreateNode(6)
//	root.RightNode.RightNode = CreateNode(7)
//	return root
//}

func (root *TreeNode) GetTreeNodeNum() int {
	if root == nil {
		return 0
	} else {
		return root.LeftNode.GetTreeNodeNum() + root.RightNode.GetTreeNodeNum() + 1
	}
}

func (root *TreeNode) GetTreeDegree() int {
	maxDegree := 0
	if root == nil {
		return maxDegree
	}
	if root.LeftNode.GetTreeDegree() > root.RightNode.GetTreeDegree() {
		maxDegree = root.LeftNode.GetTreeDegree()
	} else {
		maxDegree = root.RightNode.GetTreeDegree()
	}
	return maxDegree + 1
}

// var res []int

func (root *TreeNode) PreOrderRec() {
	if root != nil {
		// Print root
		fmt.Print(root.NodeVal, " ")
		// res = append(res, root.NodeVal)

		root.LeftNode.PreOrderRec()
		root.RightNode.PreOrderRec()
	}
}

//type Item []interface{}
//type Stack struct {
//	Items []Item
//}

//type Stack []interface{}
//
//func (stack *Stack) Push(ele interface{}) {
//	*stack = append(*stack, ele)
//}
//func (stack *Stack) Empty() bool {
//	return len(*stack) == 0
//}
//func (stack *Stack) Top() interface{} {
//	return (*stack)[len(*stack)-1]
//}
//
//func (stack *Stack) Pop() {
//	if stack.Empty() {
//		return
//	}
//	idx := len(*stack) - 1
//	*stack = (*stack)[:idx]
//}

type Stack []interface{}

// IsEmpty: check if stack is empty
func (s *Stack) empty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) push(str interface{}) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove top element of stack.
func (s *Stack) pop() {
	if s.empty() {
		return
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		*s = (*s)[:index]    // Remove it from the stack by slicing it off.
		return
	}
}
func (s *Stack) top() interface{} {
	return (*s)[len(*s)-1]
}

//func (root *TreeNode) PreOrderItr() []int {
//	var vector []int
//	var stack Stack
//	pCurrent := root
//	for !stack.empty() || pCurrent != nil {
//		if pCurrent != nil {
//			stack.push(pCurrent)
//			pCurrent = (*pCurrent).LeftNode
//		} else {
//			pNode := stack.top().(*TreeNode)
//			vector = append(vector, (*pNode).NodeVal)
//			stack.pop()
//			pCurrent = (*pNode).RightNode
//		}
//	}
//	return vector
//}

func (root *TreeNode) PreOrderItr() []int {
	var preOrder []int
	var stack Stack
	pCurrent := root
	if root == nil {
		return preOrder
	}
	stack.push(pCurrent)
	for !stack.empty() {
		pCurrent = stack.top().(*TreeNode)
		preOrder = append(preOrder, pCurrent.NodeVal)
		if pCurrent.RightNode != nil {
			stack.push(pCurrent.RightNode)
		}
		if pCurrent.LeftNode != nil {
			stack.push(pCurrent.LeftNode)
		}
	}
	//for !stack.empty() || pCurrent != nil {
	//	if pCurrent != nil {
	//		stack.push(pCurrent)
	//		pCurrent = (*pCurrent).LeftNode
	//	} else {
	//		temp := stack.top().(*TreeNode)
	//		preOrder = append(preOrder, (*temp).NodeVal)
	//		stack.pop()
	//		pCurrent = (*temp).RightNode
	//	}
	//
	//}
	return preOrder
}

func (root *TreeNode) PostOrderRec() {
	if root != nil {
		root.LeftNode.PostOrderRec()
		root.RightNode.PostOrderRec()
		// Print root
		fmt.Print(root.NodeVal, " ")
	}
}
func (root *TreeNode) PostOrderItr() {
}
func (root *TreeNode) InOrderRec() {
	if root != nil {
		root.LeftNode.InOrderRec()
		// Print root
		fmt.Print(root.NodeVal, " ")
		root.RightNode.InOrderRec()
	}
}

//func (root *TreeNode) InOrderItr() []int {
//	var inOrder []int
//	if root == nil {
//		return inOrder
//	}
//	var stack Stack
//	pCurrent := root
//	for !stack.Empty() || root != nil {
//		if pCurrent != nil {
//			stack.Push(pCurrent)
//			pCurrent = pCurrent.LeftNode
//		}
//		temp := stack.Top().(*TreeNode)
//		inOrder = append(inOrder, temp.NodeVal)
//		stack.Pop()
//		pCurrent = temp.RightNode
//	}
//	return inOrder
//}

//func (root *TreeNode) InOrderUsingStack() {
//	var stack *Stack
//	var tempNode = root
//	for tempNode != nil || !stack.empty() {
//		if tempNode != nil {
//			stack.Push(tempNode)
//			tempNode = tempNode.LeftNode
//		} else {
//			currentNode := stack.Top().(*TreeNode)
//			ans = append(ans, currentNode.NodeVal)
//			stack.Pop()
//			tempNode = tempNode.RightNode
//		}
//	}
//}

func main() {
	root := CreateNode(1)
	// left subtree
	root.LeftNode = CreateNode(2)
	root.LeftNode.LeftNode = CreateNode(4)
	root.LeftNode.RightNode = CreateNode(5)
	// right subtree
	root.RightNode = CreateNode(3)
	root.RightNode.LeftNode = CreateNode(6)
	root.RightNode.RightNode = CreateNode(7)
	root.InOrderRec()
}

//fmt.Println("Total number of nodes in trees are:", root.GetTreeNodeNum())
//fmt.Println("Depth of a binary tree is", root.GetTreeDegree())
//fmt.Println("\nTraversing tree in preorder recursively")
//root.PreOrderRec()
////fmt.Println("printing res")
////fmt.Println(res)
//fmt.Println("\nTraversing tree in postorder")
//root.PostOrderRec()
//fmt.Println("\nTraversing tree in inorder")
//root.InOrderRec()
//fmt.Println("\nTraversing tree in inOrder iteratively")
//ans := root.PreOrderItr()
//fmt.Println(ans)
