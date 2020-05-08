package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	var wait sync.WaitGroup

	wait.Add(1)

	go func() {
		ch <- "Hello world!"
		fmt.Println("Finishing goroutine...")
		wait.Done()
	}()

	time.Sleep(2 * time.Second)

	msg := <-ch
	fmt.Println(msg)
	wait.Wait()
}
