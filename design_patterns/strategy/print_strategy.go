package strategy

import "io"

type PrintStrategy interface {
	Print() error
	SetWriter(io.Writer)
	SetLog(io.Writer)
}

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (p *PrintOutput) SetWriter(w io.Writer) {
	p.Writer = w
}

func (p *PrintOutput) SetLog(w io.Writer) {
	p.LogWriter = w
}
