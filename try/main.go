package main

type Item interface{}
type Stack struct {
	Items []Item
}

func main() {
	stack := &Stack{}
	stack.Items = append(stack.Items, 2)
}
