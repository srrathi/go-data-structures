package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) append(n *node) {
	last := l.head
	for last.next != nil {
		last = last.next
	}
	last.next = n
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d\t", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteData(pos int) {
	if pos > l.length {
		return
	}
	var justPrevElement *node
	currentElement := l.head

	for i := 0; i < pos-1; i++ {
		justPrevElement = currentElement
		currentElement = currentElement.next
	}
	justPrevElement.next = currentElement.next
	l.length--
}

func (l *linkedList) deleteByValue(value int) {
	length := l.length
	if length == 0 {
		return
	}
	var justPrevElement *node
	currentElement := l.head

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	for length != 0 {
		justPrevElement = currentElement
		currentElement = currentElement.next
		if currentElement.data == value {
			l.length--
			break
		}
		if currentElement.next == nil {
			return
		}
		length--
	}

	justPrevElement.next = currentElement.next
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 75}
	node2 := &node{data: 76}
	node3 := &node{data: 77}
	node4 := &node{data: 74}
	node5 := &node{data: 73}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.append(node4)
	myList.append(node5)

	// myList.deleteData(5)
	myList.deleteByValue(100)

	myList.printListData()
}
