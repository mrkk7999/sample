package main

import "fmt"

//type Queue struct {
//	queue []int
//}
//
//func NewEmptyQueue() *Queue {
//	return &Queue{
//		queue: nil,
//	}
//}
//func NewQueue(elements []int) *Queue {
//	return &Queue{
//		queue: elements,
//	}
//}
//func (q *Queue) Enqueue(ele int) {
//	q.queue = append(q.queue, ele)
//}
//
//func (q *Queue) Dequeqe() {
//	if q.IsEmpty() {
//		return
//	}
//	q.queue = q.queue[1:]
//}
//
//func (q *Queue) Size() int {
//	return len(q.queue)
//}
//
//func (q *Queue) IsEmpty() bool {
//	return len(q.queue) == 0
//}
//
//func (q *Queue) Peek() int {
//	if q.IsEmpty() {
//		return 00
//	}
//	return q.queue[0]
//}
//func (q *Queue) PrintQueue() {
//	if q.IsEmpty() {
//		return
//	}
//	for i := len(q.queue) - 1; i >= 0; i-- {
//		fmt.Print(q.queue[i], " ")
//	}
//	fmt.Println()
//}
//func main() {
//	// var queue Queue
//	elements := []int{9, 8, 7, 6}
//	queue := NewQueue(elements)
//	//queue := NewEmptyQueue()
//	queue.Enqueue(1)
//	queue.Enqueue(2)
//	queue.Enqueue(3)
//	queue.Enqueue(4)
//	queue.PrintQueue()
//	queue.Dequeqe()
//	queue.PrintQueue()
//	currentPeek := queue.Peek()
//	fmt.Println(currentPeek)
//	fmt.Println(queue.Size())
//	fmt.Println(queue.IsEmpty())
//}

type Element []interface{}
type Queue struct {
	Elements []Element
}

func NewEmptyQueue() *Queue {
	return &Queue{
		Elements: nil,
	}
}

func NewQueue(element []Element) *Queue {
	return &Queue{
		Elements: element,
	}
}

func (queue *Queue) IsEmpty() bool {
	if len(queue.Elements) == 0 {
		return true
	}
	return false
}

// Enqueue - adds an element to the front of the queue
func (queue *Queue) Enqueue(ele Element) {
	queue.Elements = append(queue.Elements, ele)
}

// Dequeue - removes the first element of queue
func (queue *Queue) Dequeue() {
	if queue.IsEmpty() || queue == nil {
		fmt.Println("Queue is empty")
		return
	}
	queue.Elements = queue.Elements[1:len(queue.Elements)]
	return
}

// Peek - returns first element of an queue
func (queue *Queue) Peek() Element {
	if queue.IsEmpty() || queue == nil {
		fmt.Println("Queue is empty")
		return nil
	}
	return queue.Elements[0]
}

// Size - returns size of queue
func (queue *Queue) Size() int {
	if queue.IsEmpty() || queue == nil {
		fmt.Println("Queue is empty")
		return 0
	}
	return len(queue.Elements)
}

func main() {

}

//type Element []interface{}
//
//type Queue struct {
//	queue []Element
//}
//
//func NewEmptyQueue() *Queue {
//	return nil
//}
