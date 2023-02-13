package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	length := len(s.items)
	var toRemove int
	if length > 0 {
		toRemove = s.items[length-1]
		s.items = s.items[:length-1]
	}
	return toRemove
}

func main() {
	myStack := Stack{}
	myStack.Push(12)
	myStack.Push(13)
	myStack.Push(14)
	myStack.Push(15)
	myStack.Push(16)

	fmt.Println(myStack.items)

	myStack.Pop()
	fmt.Println(myStack.items)
}
