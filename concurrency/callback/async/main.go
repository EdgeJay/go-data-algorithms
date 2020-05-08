package main

import (
	"fmt"
	"strings"
	"sync"
)

func toUpperAsync(word string, f func(v string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}

var wait sync.WaitGroup

func main() {
	wait.Add(1)
	toUpperAsync("Hello callbacks!", func(v string) {
		fmt.Printf("From async callback: %s\n", v)
		wait.Done()
	})
	fmt.Println("Waiting for async response...")
	wait.Wait()
}
