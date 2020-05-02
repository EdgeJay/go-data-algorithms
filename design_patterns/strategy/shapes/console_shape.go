package shapes

import (
	"bytes"
	"io"

	"github.com/edgejay/go-data-algorithms/design_patterns/strategy"
)

type ConsoleShape struct {
	strategy.PrintOutput
}

func (c *ConsoleShape) Print() error {
	reader := bytes.NewReader([]byte("Circle\n"))
	io.Copy(c.Writer, reader)
	return nil
}
