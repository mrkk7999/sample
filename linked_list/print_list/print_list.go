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

func (head *node) printLinkedList() {
	currentNode := head
	for currentNode != nil {
		fmt.Print(currentNode.nodeVal, "\t")
		currentNode = currentNode.nextNode
	}
}
func main() {
	head := InitLinkedList()
	head.printLinkedList()
}
