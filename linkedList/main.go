package main

import "fmt"

type node struct {
	value int
	next  *node
	prev  *node
}

func (n *node) getNext() *node {
	return n.next
}

func (n *node) getPrev() *node {
	return n.prev
}

type linkedList struct {
	head *node
	tail *node
}

func (list *linkedList) first() *node {
	return list.head
}

func (list *linkedList) last() *node {
	return list.tail
}

func (list *linkedList) push(value int) {
	n := &node{value: value}
	if list.head == nil {
		list.head = n
	} else {
		n.prev = list.tail
		list.tail.next = n
	}
	list.tail = n
}

func main() {
	list := &linkedList{}
	list.push(1)
	list.push(2)
	list.push(3)

	n := list.first()
	for {
		fmt.Println(n.value)
		n = n.getNext()

		if n == nil {
			break
		}
	}

	n = list.last()
	for {
		fmt.Println(n.value)
		n = n.getPrev()

		if n == nil {
			break
		}
	}
}
