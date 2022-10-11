package main

import "fmt"

type StackInt struct {
	stk []int
}

func (s *StackInt) IsEmpty() bool {
	return len(s.stk) == 0
}
func (s *StackInt) Length() int {
	return len(s.stk)
}
func (s *StackInt) Print() {
	for i := 0; i < len(s.stk); i++ {
		fmt.Println(s.stk[i], " ")
	}
	fmt.Println()
}

func (s *StackInt) Push(ele int) {
	s.stk = append(s.stk, ele)
}
func (s *StackInt) Pop() {
	if s.IsEmpty() {
		return
	}
	s.stk = s.stk[:len(s.stk)-1]
}
func (s *StackInt) Top() int {
	return s.stk[len(s.stk)-1]
}
func main() {
	fmt.Println("Stack implementation using array")
	var stack = &StackInt{}
	stack.Push(1)
	stack.Push(2)
	stack.Print()
	currentTop := stack.Top()
	fmt.Println(currentTop)
	stack.Pop()
	fmt.Println("Stack after pop operation")
	// stack.Pop()
	stack.Print()
	fmt.Printf("Currrent length of stack is: %d", stack.Length())
}
