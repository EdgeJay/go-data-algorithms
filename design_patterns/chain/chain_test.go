package chain

import (
	"fmt"
	"strings"
	"testing"
)

type myTestWriter struct {
	receivedMessage *string
}

func (w *myTestWriter) Write(p []byte) (int, error) {
	if w.receivedMessage == nil {
		w.receivedMessage = new(string)
	}

	temp := fmt.Sprintf("%s%s", *w.receivedMessage, p)
	w.receivedMessage = &temp
	return len(p), nil
}

func (w *myTestWriter) Next(s string) {
	w.Write([]byte(s))
}

func TestCreateChain(t *testing.T) {
	t.Log(`3 loggers, 2 of them writes to console, 2nd logger only writes to console if it found the world "hello"`)
	t.Log(`3rd logger writes to some variable if 2nd logger found "hello"`)

	myWriter := myTestWriter{}
	writerLogger := WriterLogger{Writer: &myWriter}
	secondLogger := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &secondLogger}

	t.Run("WriterLogger should not receive any message", func(t *testing.T) {
		chain.Next("message that breaks the chain")

		if myWriter.receivedMessage != nil {
			t.Errorf("Expected WriterLogger to not receive message, got %s", *myWriter.receivedMessage)
		}
	})

	t.Run("WriterLogger should receive message", func(t *testing.T) {
		chain.Next("hello")

		if myWriter.receivedMessage == nil || !strings.Contains(*myWriter.receivedMessage, "hello") {
			t.Error("Expected WriterLogger to receive message, got nil")
		}
	})
}
