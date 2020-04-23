package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int) {
	n := 0
	for {
		fmt.Printf("-> Send job: %d\n", n)
		ch <- n
		n++
	}
}

func echoWorker(in, out chan int) {
	for {
		n := <-in
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		out <- n
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	for i := 0; i < 4; i++ {
		go echoWorker(in, out)
	}
	go producer(in)

	for n := range out {
		fmt.Printf("Received job: %d\n", n)
	}
}
