package bridge

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestPrinterImpl1(t *testing.T) {
	printer := printerImpl1{}
	err := printer.printMessage("foo")

	if err != nil {
		t.Fatalf("Expected error to be nil, got error: %s", err.Error())
	}
}

type testWriter struct {
	msg string
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		w.msg = string(p)
	} else {
		err = errors.New("Content received on writer was empty")
	}
	return
}

func TestPrinterImpl2(t *testing.T) {
	printer := printerImpl2{}
	err := printer.printMessage("foo")
	expectedError := "Missing io.Writer"

	if err != nil && !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected error to be \"%s\", got error: \"%s\"", expectedError, err.Error())
	}

	writer := testWriter{}
	printer = printerImpl2{
		writer: &writer,
	}

	expectedMessage := "Hello"
	err = printer.printMessage(expectedMessage)

	if err != nil {
		t.Errorf("Expected error to be nil, got error: \"%s\"", err.Error())
	}

	if writer.msg != expectedMessage {
		t.Errorf("Expected writer.msg to be \"%s\", got \"%s\"", expectedMessage, writer.msg)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello printer"
	nPrinter := normalPrinter{
		msg:     expectedMessage,
		printer: &printerImpl1{},
	}

	err := nPrinter.print()
	if err != nil {
		t.Errorf("Expected error to be nil, got %s", err.Error())
	}

	writer := testWriter{}
	nPrinter = normalPrinter{
		msg: expectedMessage,
		printer: &printerImpl2{
			writer: &writer,
		},
	}

	err = nPrinter.print()
	if err != nil {
		t.Errorf("Expected error to be nil, got %s", err.Error())
	}

	if writer.msg != expectedMessage {
		t.Errorf("Expected writer.msg to be \"%s\", got \"%s\"", expectedMessage, writer.msg)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello printer"
	expectedMessage := fmt.Sprintf("Message from Packt: %s", passedMessage)
	pktPrinter := packtPrinter{
		msg:     passedMessage,
		printer: &printerImpl1{},
	}

	err := pktPrinter.print()
	if err != nil {
		t.Errorf("Expected error to be nil, got %s", err.Error())
	}

	writer := testWriter{}
	pktPrinter = packtPrinter{
		msg: passedMessage,
		printer: &printerImpl2{
			writer: &writer,
		},
	}

	err = pktPrinter.print()
	if err != nil {
		t.Errorf("Expected error to be nil, got %s", err.Error())
	}

	if writer.msg != expectedMessage {
		t.Errorf("Expected writer.msg to be \"%s\", got \"%s\"", expectedMessage, writer.msg)
	}
}
