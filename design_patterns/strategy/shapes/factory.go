package shapes

import (
	"fmt"
	"os"

	"github.com/edgejay/go-data-algorithms/design_patterns/strategy"
)

const (
	ConsoleStrategy = "console"
	ImageStrategy   = "image"
)

func NewPrinter(s string) (strategy.PrintStrategy, error) {
	switch s {
	case ConsoleStrategy:
		return &ConsoleShape{
			PrintOutput: strategy.PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case ImageStrategy:
		return &ImageShape{
			PrintOutput: strategy.PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy %s not found\n", s)
	}
}
