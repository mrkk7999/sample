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

// type Element []interface{}

type Queue struct {
	Elements []interface{}
}

func NewEmptyQueue() *Queue {
	return &Queue{
		Elements: nil,
	}
}
func (q *Queue) Enqueue(ele interface{}) {
	q.Elements = append(q.Elements, ele)
}
func (q *Queue) IsEmpty() bool {
	if len(q.Elements) == 0 {
		return true
	}
	return false
}
func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		return
	}
	q.Elements = q.Elements[1:]
}
func (q *Queue) Peak() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.Elements[0]
}

func (q *Queue) Size() int {
	if q.IsEmpty() {
		return 0
	}
	return len(q.Elements)
}

//func (root *TreeNode) LevelOrder() [][]int {
//	var res [][]int
//	if root == nil {
//		return res
//	}
//	queue := NewEmptyQueue()
//	queue.Enqueue(root)
//	for !queue.IsEmpty() {
//		size := queue.Size()
//		level := []int{}
//		for i := 0; i < size; i++ {
//			currentNode := queue.Peak().(*TreeNode)
//			queue.Dequeue()
//			if currentNode.LeftNode != nil {
//				queue.Enqueue(currentNode.LeftNode)
//			}
//			if currentNode.RightNode != nil {
//				queue.Enqueue(currentNode.RightNode)
//			}
//			level = append(level, currentNode.NodeVal)
//		}
//		res = append(res, level)
//	}
//	return res
//}

func (root *TreeNode) LevelOrder() [][]int {
	var res [][]int
	queue := NewEmptyQueue()
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		size := queue.Size()
		var level []int
		for i := 0; i < size; i++ {
			// we'll take front element of queue
			currentNode := queue.Peak().(*TreeNode)
			// to remove it from front of queue
			// as it's operations is over we'll head to next element
			queue.Dequeue()
			// we'll insert left element from back(i.e enqueue)
			// which will ultimately our first element
			// when we read from front(i.e peak)
			if currentNode.LeftNode != nil {
				queue.Enqueue(currentNode.LeftNode)
			}
			// we'll insert right node from back after left as left will be taken first
			// and then our right as required in level order traversal
			if currentNode.RightNode != nil {
				queue.Enqueue(currentNode.RightNode)
			}
			// we continue fill it till as our queue doesn't get empty
			// and at each iteration queue is only getting filled with
			level = append(level, currentNode.NodeVal)
		}
		res = append(res, level)
	}
	return res
}
func main() {
	root := InitTree()
	res := root.LevelOrder()
	fmt.Println(res)
}
