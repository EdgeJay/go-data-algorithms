package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	numRoutines := 5
	wait.Add(numRoutines)
	for index := 0; index < numRoutines; index++ {
		go func(index int) {
			fmt.Printf("Hello world from routine: %d\n", index)
			wait.Done()
		}(index)
	}
	wait.Wait()
}
