package factory

import (
	"fmt"
	"testing"
)

func TestGetPaymentMethodName(t *testing.T) {
	paymentMethod := getPaymentMethod(cash)

	if paymentMethod.getName() != "cash" {
		t.Errorf("Expected payment method name to be \"cash\", got \"%s\"", paymentMethod.getName())
	}
}

func TestGetPaymentMethodMessage(t *testing.T) {
	expectedAmount := 32.
	expectedMessage := fmt.Sprintf("Paid $%.2f using cash payment method", expectedAmount)
	paymentMethod := getPaymentMethod(cash)
	message, err := paymentMethod.pay(expectedAmount)

	if err != nil {
		t.Errorf("Expected pay method to return nil error, got %#v", err)
	}

	if message != expectedMessage {
		t.Errorf("Expected message to be \"%s\", got \"%s\"", expectedMessage, message)
	}
}

func TestGetPaymentMethodError(t *testing.T) {
	amount := -10.
	paymentMethod := getPaymentMethod(cash)
	message, err := paymentMethod.pay(amount)
	expectedMessage := fmt.Sprintf("Invalid amount: %.2f", amount)

	if message != "" {
		t.Errorf("Expected message to be \"\", got \"%s\"", message)
	}

	if err == nil {
		t.Fatal("Expected pay method to return error, got nil")
	}

	if actualMessage := fmt.Sprint(err); actualMessage != expectedMessage {
		t.Errorf("Expected pay method to return error with message \"%s\", got \"%s\"", expectedMessage, actualMessage)
	}
}

func TestGetPaymentMethodForCreditCard(t *testing.T) {
	paymentMethod := getPaymentMethod(creditCard)

	if paymentMethod.getName() != "credit card" {
		t.Errorf("Expected payment method name to be \"credit card\", got \"%s\"", paymentMethod.getName())
	}

	expectedAmount := 45.
	expectedMessage := fmt.Sprintf("Paid $%.2f using credit card payment method", expectedAmount)
	message, err := paymentMethod.pay(expectedAmount)

	if err != nil {
		t.Errorf("Expected pay method to return nil error, got %#v", err)
	}

	if message != expectedMessage {
		t.Errorf("Expected message to be \"%s\", got \"%s\"", expectedMessage, message)
	}
}

func TestGetPaymentMethodNil(t *testing.T) {
	paymentMethod := getPaymentMethod(3)

	if paymentMethod != nil {
		t.Error("Expected payment method to be nil when invalid method is provided")
	}
}
