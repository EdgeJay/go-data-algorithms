package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
