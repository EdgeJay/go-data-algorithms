package main

import (
	"fmt"
	"time"
)

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			fmt.Println(msg)
		case msg := <-goodbyeCh:
			fmt.Println(msg)
		case <-time.After(time.Second * 2):
			fmt.Println("Nothing received in 2 seconds, exiting...")
			quitCh <- true
			break
		}
	}
}

func main() {
	helloCh := make(chan string)
	goodbyeCh := make(chan string)
	quitCh := make(chan bool)

	go receiver(helloCh, goodbyeCh, quitCh)
	go sendString(helloCh, "hello!")

	time.Sleep(time.Second)

	go sendString(goodbyeCh, "goodbye!")

	time.Sleep(time.Second * 5)

	<-quitCh
}
