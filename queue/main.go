package main

import (
	"errors"
	"fmt"
)

type queue struct {
	items []int
}

func (q *queue) enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *queue) dequeue() (int, error) {
	if len(q.items) == 0 {
		return -1, errors.New("Queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func main() {
	q := &queue{}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)

	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}
