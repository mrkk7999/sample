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
func main() {
	head := InitLinkedList()
	currentNode := head
	for currentNode.nextNode != nil {
		fmt.Print(currentNode.nodeVal, "\t")
		currentNode = currentNode.nextNode
	}
	fmt.Print(currentNode.nodeVal)
}
