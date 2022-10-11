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
	head = initWithHead(1)
	tail := head
	for i := 2; i <= 10; i++ {
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

func (head *node) removeFromBack(n int) *node {
	//dummy:=&node{}
	//dummy.nextNode = head
	cnt := head
	N := 0
	for cnt != nil {
		cnt = cnt.nextNode
		N = N + 1
	}
	pos := N - n
	if pos == 0 {
		return head.nextNode
	} else if pos < 0 {
		return head
	}
	prev := head
	for pos > 1 {
		prev = prev.nextNode
		pos = pos - 1
	}
	prev.nextNode = prev.nextNode.nextNode
	// 1 2 3 4 5 6
	// N = 6
	// n = 4
	// pos = 2
	// iteration
	// 1st pos>0 -> true, prev at 2, pos = 1
	// 2nd pos>0 -> true , prev at 3 pos = 0
	return head
}

//	func (head *node) temp() {
//		t := head
//		pos := 2
//		for pos > 1 {
//			t = t.nextNode
//			pos = pos - 1
//		}
//		fmt.Println(t.nodeVal)
//	}
func main() {
	head := InitLinkedList()
	head.printLinkedList()
	head = head.removeFromBack(66)
	fmt.Println()
	head.printLinkedList()
}
