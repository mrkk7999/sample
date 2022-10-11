package main

func (root *TreeNode) postOrderUsingOneStack() []int {
	var current = root
	var stack Stack
	var res []int
	for current != nil || !stack.isEmpty() {
		if current != nil {
			stack.push(current)
			current = current.LeftNode
		} else {
			temp := stack.top().(*TreeNode).RightNode
			if temp == nil {
				temp = stack.top().(*TreeNode)
				stack.pop()
				res = append(res, temp.NodeVal)
				for !stack.isEmpty() && temp == stack.top().(*TreeNode).RightNode {
					temp = stack.top().(*TreeNode)
					stack.pop()
					res = append(res, temp.NodeVal)
				}
			} else {
				current = temp

			}
		}
	}
	return res
}
