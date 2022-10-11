package main

import (
	"fmt"
)

// https://medium.com/faun/linked-list-in-golang-94aff9230321

// Structure of node

type Node struct {
	NodeVal  interface{}
	NextNode *Node
}

type LinkedList struct {
	LenOfLL int
	Head    *Node
}

func InitEmptyList() *LinkedList {
	return &LinkedList{}
}
func InitWithHead(Name string) *LinkedList {
	return &LinkedList{
		LenOfLL: 1,
		Head: &Node{
			NodeVal:  Name,
			NextNode: nil,
		},
	}
}
func (list *LinkedList) AddFront(Name string) {
	node := &Node{
		NodeVal:  Name,
		NextNode: nil,
	}
	if list.Head == nil {
		list.Head = node
	} else {
		node.NextNode = list.Head
		list.Head = node
	}
	list.LenOfLL++
	return
}

func (list *LinkedList) AddBack(Name string) {
	node := &Node{
		NodeVal:  Name,
		NextNode: nil,
	}
	if list.Head == nil {
		list.Head = node
	} else {
		current := list.Head
		for current.NextNode != nil {
			current = current.NextNode
		}
		current.NextNode = node
	}
	list.LenOfLL++
	return
}

func (list *LinkedList) Traverse() error {
	if list.Head == nil {
		return fmt.Errorf("TraverseError :  List is empty")
	}
	current := list.Head
	for current != nil {
		fmt.Print(current.NodeVal, " ")
		current = current.NextNode
	}
	return nil
}

func (list *LinkedList) RemoveFirst() error {
	if list.Head == nil {
		return fmt.Errorf("empty list")
	}
	list.Head = list.Head.NextNode
	list.LenOfLL--
	return nil
}

func (list *LinkedList) RemoveBack() {
	if list.Head == nil {
		return
	}
	var prev *Node
	current := list.Head
	for current.NextNode != nil {
		prev = current
		current = current.NextNode
	}
	if prev != nil {
		prev.NextNode = nil
	} else {
		list.Head = nil
	}
	list.LenOfLL--
	return
}
func main() {
	singleList := InitEmptyList()
	singleList = InitWithHead("MyName:\t")
	singleList.AddBack("Kiran")
	singleList.AddBack("Arun")
	singleList.AddBack("Kshirsagar")
	singleList.Traverse()
	singleList.RemoveFirst()
	fmt.Println()
	singleList.Traverse()
	singleList.RemoveBack()
	fmt.Println()
	singleList.Traverse()
}
