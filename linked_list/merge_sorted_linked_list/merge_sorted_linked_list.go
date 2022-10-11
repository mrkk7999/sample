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
func InitLinkedListOne() *node {
	head := initEmptyList()
	head = initWithHead(0)
	tail := head
	for i := 2; i < 10; i = i + 2 {
		tail.nextNode = initWithHead(i)
		tail = tail.nextNode
	}
	return head
}
func InitLinkedListTwo() *node {
	head := initEmptyList()
	head = initWithHead(1)
	tail := head
	for i := 3; i < 10; i = i + 2 {
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
	fmt.Println()
}

func mergeTwoLists(l1 *node, l2 *node) *node {
	mL := &node{}
	mL.nextNode = nil
	for l1 != nil || l2 != nil {
		if l1 == nil {
			if mL == nil {
				mL = l2
			} else {
				mL.nextNode = l2
			}
		}
		if l2 == nil {
			if mL == nil {
				mL = l2
			} else {
				mL.nextNode = l1
			}
		}
		if l1.nodeVal.(int) <= l2.nodeVal.(int) {
			if mL == nil {
				mL = l1
				l1 = l1.nextNode
			} else {
				mL.nextNode = l1
				l1 = l1.nextNode
			}
		} else {
			if mL == nil {
				mL = l1
				l1 = l1.nextNode
			} else {
				mL.nextNode = l1
				l1 = l1.nextNode
			}
		}
		if l1.nodeVal.(int) < l2.nodeVal.(int) {
			mL.nodeVal = l1.nodeVal
			l1 = l1.nextNode
		} else {
			mL.nodeVal = l2.nodeVal
			l2 = l2.nextNode
		}
	}

	return mL
}

func main() {
	headL1 := InitLinkedListOne()
	headL2 := InitLinkedListTwo()
	headL1.printLinkedList()
	headL2.printLinkedList()
	mergedList := mergeTwoLists(headL1, headL2)
	mergedList.printLinkedList()
}
