package main

import "fmt"

type Element interface{}

type Stack struct {
	Elements []Element
}

func NewEmptyStack() *Stack {
	return &Stack{
		Elements: nil,
	}
}

func NewStack(element []Element) *Stack {
	fmt.Println(element)
	//return &Stack{
	//	element: element,
	//}
	tempStack := &Stack{}
	for i := 0; i < len(element); i++ {
		tempStack.Elements = append(tempStack.Elements, element[i])
	}
	return tempStack
}

func (stack *Stack) Push(element Element) {
	fmt.Println(element)
	stack.Elements = append(stack.Elements, element)
}
func (stack *Stack) IsEmpty() bool {
	if stack == nil || len(stack.Elements) == 0 {
		return true
	}
	return false
}
func (stack *Stack) Size() int {
	return len(stack.Elements)
}
func (stack *Stack) PrintStack() {
	if stack.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println()
	for i := 0; i < stack.Size(); i++ {
		fmt.Print(stack.Elements[i], " ")
	}
}
func (stack *Stack) Pop() {
	if stack.IsEmpty() {
		return
	}
	stack.Elements = stack.Elements[:len(stack.Elements)-1]
}

func (stack *Stack) Top() Element {
	if stack.IsEmpty() {
		return nil
	}
	return stack.Elements[stack.Size()-1]
}
func main() {
	fmt.Println("Creating stack")
	stack := NewEmptyStack()
	fmt.Println("Pushing value 1 and 2 on stack")
	stack.Push(1)
	stack.PrintStack()
	//arr := []Element{2, 3, 4, 5}
	//stack = NewStack(arr)
	//stack.PrintStack()
	//stack.Push(6)
	//stack.PrintStack()
	stack.Push(2)
	stack.PrintStack()
	stack.Size()
	stack.PrintStack()
	fmt.Println("\n Removing the top element")
	stack.Pop()
	fmt.Println("\n Stack after pop operation")
	stack.PrintStack()
	fmt.Println("\n Printing the top element")
	fmt.Println(stack.Top())
}
