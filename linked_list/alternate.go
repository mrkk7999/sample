package main

type Node struct {
	Data     int
	NextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList LinkedList) AddToHead(data int) {
	var tempNode = Node{}
	tempNode.Data = data
	if tempNode.NextNode != nil {
		tempNode.NextNode = linkedList.headNode
	}
	linkedList.headNode = &tempNode
}
