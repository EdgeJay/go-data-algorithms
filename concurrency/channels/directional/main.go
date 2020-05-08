package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func(ch chan<- string) {
		ch <- "Hello World!"
		fmt.Println("Finishing goroutine...")
	}(ch)

	time.Sleep(time.Second)

	receivingCh(ch)
}

func receivingCh(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)
}
