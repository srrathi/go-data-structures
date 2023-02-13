package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue - add element in a queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue - remove first element from queue
func (q *Queue) Dequeue() int {
	length := len(q.items)
	var itemToRemove int
	if length > 0 {
		itemToRemove = q.items[0]
		q.items = q.items[1:]
	}
	return itemToRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.Enqueue(4)
	myQueue.Enqueue(5)
	myQueue.Enqueue(6)

	fmt.Println(myQueue.items)
	myQueue.Dequeue()
	fmt.Println(myQueue.items)
}
