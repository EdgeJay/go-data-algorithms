package factory

import (
	"fmt"
)

const (
	cash       = 1
	creditCard = 2
)

type paymentMethod interface {
	getName() string
	pay(amount float64) (message string, err error)
}

type cashPaymentMethod struct {
	name          string
	messageFormat string
}

func (m *cashPaymentMethod) getName() string {
	return "cash"
}

func (m *cashPaymentMethod) pay(amount float64) (message string, err error) {
	if amount >= 0. {
		message = fmt.Sprintf(m.messageFormat, amount)
	} else {
		err = fmt.Errorf("Invalid amount: %.2f", amount)
	}
	return
}

type creditCardPaymentMethod struct {
	name          string
	messageFormat string
}

func (m *creditCardPaymentMethod) getName() string {
	return "credit card"
}

func (m *creditCardPaymentMethod) pay(amount float64) (message string, err error) {
	if amount >= 0. {
		message = fmt.Sprintf(m.messageFormat, amount)
	} else {
		err = fmt.Errorf("Invalid amount: %.2f", amount)
	}
	return
}

func getPaymentMethod(method int) paymentMethod {
	switch method {
	case cash:
		return &cashPaymentMethod{messageFormat: "Paid $%.2f using cash payment method"}
	case creditCard:
		return &creditCardPaymentMethod{messageFormat: "Paid $%.2f using credit card payment method"}
	}
	return nil
}
