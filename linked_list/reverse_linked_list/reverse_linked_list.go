package main

import "fmt"

// Structure
// Node -> individual node of linked list
type node struct {
	nodeVal  interface{}
	nextNode *node
}

func initEmptyHead() *node {
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
	head := initEmptyHead()
	head = initWithHead(0)
	tail := head
	for i := 1; i < 10; i++ {
		tail.nextNode = initWithHead(i)
		tail = tail.nextNode
	}
	return head
}
func (head *node) printLinkedList() {
	temp := head
	for temp != nil {
		fmt.Print(temp.nodeVal, "\t")
		temp = temp.nextNode
	}
	fmt.Println()
}

func (head *node) reverseLinkedList() *node {
	var dummyNode *node
	for head != nil {
		next := head.nextNode
		head.nextNode = dummyNode
		dummyNode = head
		head = next
	}
	return dummyNode
}

func main() {
	head := InitLinkedList()
	currentNode := head
	for true {
		if currentNode == nil {
			break
		}
		fmt.Print(currentNode.nodeVal, "\t")
		currentNode = currentNode.nextNode
	}
	head.printLinkedList()
	fmt.Println()
	head = head.reverseLinkedList()
	head.printLinkedList()
}
