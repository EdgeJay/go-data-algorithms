package main

import "fmt"

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Printf("> Executing command: %s\n", c.message)
}

func CreateCommand(s string) Command {
	fmt.Printf("Creating command: %s\n", s)
	return &ConsoleOutput{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (q *CommandQueue) AddCommand(c Command) {
	q.queue = append(q.queue, c)
	if len(q.queue) == 3 {
		for _, command := range q.queue {
			command.Execute()
		}
		q.queue = make([]Command, 3)
	}
}

func main() {
	queue := CommandQueue{}
	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
}
