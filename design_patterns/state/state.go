package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type GameState interface {
	executeState(*GameContext) bool
}

type GameContext struct {
	UserAnswer   int
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}

type StartState struct{}

func (s *StartState) executeState(c *GameContext) bool {
	c.Next = &AskState{}

	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)

	fmt.Println("Set number of retries for game difficulty")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)

	return true
}

type ErrorState struct{}

func (e *ErrorState) executeState(c *GameContext) bool {
	fmt.Printf("%d is not a valid number\n", c.UserAnswer)
	c.Next = &AskState{}
	return true
}

type FinishState struct{}

func (f *FinishState) executeState(c *GameContext) bool {
	if c.Won {
		c.Next = &WinState{}
	} else {
		c.Next = &LoseState{}
	}

	return true
}

type AskState struct{}

func (a *AskState) executeState(c *GameContext) bool {
	fmt.Printf("Guess a number between 0 and 10, you have %d retries left.\n", c.Retries)

	var n int
	fmt.Fscanf(os.Stdin, "%d\n", &n)

	c.Retries--
	c.UserAnswer = n

	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}

	if n < 0 {
		c.Next = &ErrorState{}
	}

	if c.Retries == 0 {
		c.Next = &FinishState{}
	}

	return true
}

type WinState struct{}

func (w *WinState) executeState(c *GameContext) bool {
	fmt.Println("Congratulations, you won!")
	return false
}

type LoseState struct{}

func (l *LoseState) executeState(c *GameContext) bool {
	fmt.Printf("You lose! The correct answer is %d.\n", c.SecretNumber)
	return false
}

func main() {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}

	for game.Next.executeState(&game) {
	}
}
