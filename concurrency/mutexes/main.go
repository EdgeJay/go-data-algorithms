package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}

	for index := 0; index < 10; index++ {
		go func(index int) {
			counter.Lock()
			defer counter.Unlock()
			counter.value++
			fmt.Printf("Counter new value: %d\n", counter.value)
		}(index)
	}

	time.Sleep(time.Second)

	counter.Lock()
	defer counter.Unlock()

	fmt.Println(counter.value)
}
