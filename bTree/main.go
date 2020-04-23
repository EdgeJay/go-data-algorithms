package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func (n *node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func (n *node) exists(value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	}

	if n.value > value {
		return n.left.exists(value)
	}

	return n.right.exists(value)
}

type tree struct {
	root *node
}

func (t *tree) insert(value int) *tree {
	if t.root == nil {
		t.root = &node{value: value}
	} else {
		t.root.insert(value)
	}
	return t
}

func printNode(n *node) {
	if n == nil {
		return
	}

	fmt.Println("Node value: ", n.value)
	printNode(n.left)
	printNode(n.right)
}

func main() {
	t := &tree{}
	t.insert(10).
		insert(8).
		insert(20).
		insert(9).
		insert(0).
		insert(15).
		insert(25)
	printNode(t.root)

	fmt.Println(t.root.exists(8))
	fmt.Println(t.root.exists(11))
	fmt.Println(t.root.exists(25))

	t.insert(11)
	fmt.Println(t.root.exists(11))
}
