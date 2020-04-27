package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &pizzaDecorator{}
	pizzaResult, _ := pizza.addIngredient()
	expectedText := "Pizza with following ingredients:"

	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("Expected addIngredient to return \"%s\", got \"%s\"", expectedText, pizzaResult)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	testOnion := &onion{}
	onionResult, err := testOnion.addIngredient()
	expectedText := "onions"

	if err == nil {
		t.Errorf("Expected error to be returned if addIngredient is called when ingredient field is nil, got \"%s\"", onionResult)
	}

	testOnion = &onion{&pizzaDecorator{}}
	onionResult, err = testOnion.addIngredient()

	if err != nil {
		t.Errorf("Expected error to be nil, got \"%s\"", err.Error())
	}

	if !strings.Contains(onionResult, expectedText) {
		t.Errorf("Expected result to contain \"%s\", got \"%s\"", expectedText, onionResult)
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	testMeat := &meat{}
	meatResult, err := testMeat.addIngredient()
	expectedText := "meat"

	if err == nil {
		t.Errorf("Expected error to be returned if addIngredient is called when ingredient field is nil, got \"%s\"", meatResult)
	}

	testMeat = &meat{&pizzaDecorator{}}
	meatResult, err = testMeat.addIngredient()

	if err != nil {
		t.Errorf("Expected error to be nil, got \"%s\"", err.Error())
	}

	if !strings.Contains(meatResult, expectedText) {
		t.Errorf("Expected result to contain \"%s\", got \"%s\"", expectedText, meatResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &onion{&meat{&pizzaDecorator{}}}
	result, err := pizza.addIngredient()
	expectedText := "Pizza with following ingredients: meat onions"

	if err != nil {
		t.Errorf("Expected error to be nil, got \"%s\"", err.Error())
	}

	if !strings.Contains(result, expectedText) {
		t.Errorf("Expected result to contain \"%s\", got \"%s\"", expectedText, result)
	}
}
