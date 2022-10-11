package main

type queue struct {
	elements []int
}

func (q *queue) enqueue(newEle int) {
	q.elements = append(q.elements, newEle)
}
func (q *queue) dequeue() {

}
func main() {

}
