package chain

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(str string) {
	fmt.Printf("First logger: %s\n", str)

	if f.NextChain != nil {
		f.NextChain.Next(str)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (s *SecondLogger) Next(str string) {
	if strings.Contains(strings.ToLower(str), "hello") {
		fmt.Printf("Second logger: %s\n", str)

		if s.NextChain != nil {
			s.NextChain.Next(str)
		}

		return
	}

	fmt.Println("Finished in second logging")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(str string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + str))
	}

	if w.NextChain != nil {
		w.NextChain.Next(str)
	}
}
