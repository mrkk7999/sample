package main

func (root *TreeNode) postOrder2Stack() []int {
	var stkOne, stkTwo Stack
	var res []int
	var currentNode = root
	stkOne.push(root)
	for !stkOne.isEmpty() {
		currentNode = stkOne.top().(*TreeNode)
		stkOne.pop()
		stkTwo.push(currentNode.NodeVal)
		if currentNode.LeftNode != nil {
			stkOne.push(currentNode.LeftNode)
		}
		if currentNode.RightNode != nil {
			stkOne.push(currentNode.RightNode)
		}
	}

	for !stkTwo.isEmpty() {
		temp := stkTwo.top().(int)
		res = append(res, temp)
		stkTwo.pop()
	}
	return res
}
