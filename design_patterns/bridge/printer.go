package bridge

import (
	"errors"
	"fmt"
	"io"
)

type printerAPI interface {
	printMessage(msg string) error
}

type printerImpl1 struct{}

func (p *printerImpl1) printMessage(msg string) error {
	fmt.Println(msg)
	return nil
}

type printerImpl2 struct {
	writer io.Writer
}

func (p *printerImpl2) printMessage(msg string) error {
	if p.writer == nil {
		return errors.New("Missing io.Writer")
	}

	_, err := fmt.Fprintf(p.writer, "%s", msg)
	return err
}

type printerAbstraction interface {
	print() error
}

type normalPrinter struct {
	msg     string
	printer printerAPI
}

func (p *normalPrinter) print() error {
	err := p.printer.printMessage(p.msg)
	return err
}

type packtPrinter struct {
	msg     string
	printer printerAPI
}

func (p *packtPrinter) print() error {
	err := p.printer.printMessage(fmt.Sprintf("Message from Packt: %s", p.msg))
	return err
}
