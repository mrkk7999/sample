package main

import "fmt"

// Structure
// Node -> individual node of linked list
type node struct {
	nodeVal  interface{}
	nextNode *node
}

func initEmptyList() *node {
	return &node{
		nodeVal:  nil,
		nextNode: nil,
	}
}

func initWithHead(ele interface{}) *node {
	return &node{
		nodeVal:  ele,
		nextNode: nil,
	}
}
func InitLinkedList() *node {
	head := initEmptyList()
	head = initWithHead(0)
	tail := head
	for i := 1; i < 10; i++ {
		tail.nextNode = initWithHead(i)
		tail = tail.nextNode
	}
	return head
}
func (root *node) printLinkedList() {
	currentNode := root
	for currentNode != nil {
		fmt.Print(currentNode.nodeVal, "\t")
		currentNode = currentNode.nextNode
	}
}

func (root *node) addToTail(val int) *node {
	tempNode := &node{nodeVal: val}
	tail := root
	for tail.nextNode != nil {
		tail = tail.nextNode
	}
	tail.nextNode = tempNode
	return root
}

func main() {
	// initial list
	//  0       1       2       3       4       5       6       7       8       9
	head := InitLinkedList()
	head = head.addToTail(79)
	head.printLinkedList()
}
