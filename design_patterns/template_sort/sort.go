package main

import (
	"fmt"
	"sort"
)

type myList []int

func (list myList) Len() int {
	return len(list)
}

func (list myList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list myList) Less(i, j int) bool {
	return list[i] < list[j]
}

func main() {
	var list myList = []int{6, 4, 2, 8, 1}
	fmt.Println(list)
	sort.Sort(list)
	fmt.Println(list)
}
