package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)
	go func() {
		ch <- "Message 1: Hello world!"
		ch <- "Message 2: Hello world!!"
		fmt.Println("Finishing goroutine...")
	}()

	time.Sleep(time.Second)

	msg := <-ch
	fmt.Println(msg)
}
