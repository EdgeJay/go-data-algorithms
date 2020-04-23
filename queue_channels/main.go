package main

import (
	"fmt"
)

type queue struct {
	items chan int
}

func (q *queue) enqueue(value int) {
	q.items <- value
}

func (q *queue) dequeue() int {
	return <-q.items
}

func main() {
	q := &queue{items: make(chan int, 16)}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)

	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}
