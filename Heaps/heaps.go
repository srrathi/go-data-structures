package main

import "fmt"

// HEAPS ARE VISUALISED AS TREES BUT ARE STOERD AS AN ARRAY.
// THERE ARE MAX HEAP AND MIN HEAP.
// IN MAX HEAP EACH PARENT IS GREATER THAN ANY OF ITS CHILD ELEMENT.
// IN MIN HEAP EACH PARENT IS LESS THAN ANY OF ITS CHILD ELEMENT.
// INDEX_OF_LEFT_CHILD = INDEX_OF_PARENT*2 + 1 (ALWAYS ODD)
// INDEX_OF_RIGHT_CHILD = INDEX_OF_PARENT*2 + 2
// INDEX_OF_PARENT = (INDEX_OF_LEFT_CHILD - 1) / 2

// MAX HEAP STRUCT IS A SLICE
type MaxHeap struct {
	array []int
}

// INSERT ADDS AN ELEMENT IN THE HEAP
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// EXTRACT RETURNS THE LARGEST KEY FROM THE HEAP AND REMOVE IT FROM THE HEAP
func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}
	extracted := h.array[0]
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// MAX HEAPIFY WILL HEAPIFY FROM BOTTOM TO TOP
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = 1
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1 int, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	m.Insert(10)
	m.Insert(15)
	m.Insert(19)
	m.Insert(25)
	m.Insert(38)
	m.Insert(50)
	fmt.Println(m)
}
