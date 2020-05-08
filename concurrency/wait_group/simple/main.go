package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		fmt.Println("Hello world!")
		wait.Done()
		// wait.Add(-1) // since Add method accepts a delta, -1 works the same as wait.Done()
	}()
	wait.Wait()
}
