package main

import (
	"fmt"
	"sync"

	"github.com/edgejay/go-data-algorithms/design_patterns/singleton"
)

var lock sync.Mutex

func addOneInThread(ch chan<- int) {
	lock.Lock()
	defer lock.Unlock()

	singletonCounter := singleton.GetInstance()

	singletonCounter.AddOne()
	ch <- singletonCounter.Value()
}

func main() {
	singletonCounter := singleton.GetInstance()
	fmt.Println("singletonCounter.AddOne() : ", singletonCounter.AddOne())
	fmt.Println("singletonCounter.AddOne() : ", singletonCounter.AddOne())
	fmt.Println("singletonCounter.AddOne() : ", singletonCounter.AddOne())

	anotherCounter := singleton.GetInstance()
	fmt.Println("anotherCounter.Value() : ", anotherCounter.Value())

	ch := make(chan int)

	go addOneInThread(ch)
	go addOneInThread(ch)
	go addOneInThread(ch)

	fmt.Printf("counter value: %d\n", <-ch)
	fmt.Printf("counter value: %d\n", <-ch)
	fmt.Printf("counter value: %d\n", <-ch)
}
