package main

import "fmt"

// FIFO
type Queue struct{
	element []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (e *Queue)GetQueueLength() int {
	return len(e.element)
}

func (e *Queue)IsEmpty() bool {
	if len(e.element) == 0 {
		return true
	} else {
		return false
	}
}

func (e *Queue)Push(num int)  {
	e.element = append(e.element, num)
}

func (e *Queue)Pop() bool {
	if e.IsEmpty() {
		return false
	} else {
		e.element = e.element[1:]
		return true
	}
}

func (e *Queue)Front() int {
	if e.IsEmpty() {
		return -1
	} else {
		return e.element[0]
	}
}
func (e *Queue) Print()  {
	for i:=0; i<e.GetQueueLength(); i++ {
		fmt.Printf("%d ", e.element[i])
	}
	fmt.Println()
	fmt.Printf("front: %d\n", e.Front())
}

