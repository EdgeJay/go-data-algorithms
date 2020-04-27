package decorator

import (
	"errors"
	"fmt"
)

type ingredientAdd interface {
	addIngredient() (string, error)
}

type pizzaDecorator struct {
	ingredient ingredientAdd
}

func (p *pizzaDecorator) addIngredient() (string, error) {
	return "Pizza with following ingredients:", nil
}

type meat struct {
	ingredient ingredientAdd
}

func (m *meat) addIngredient() (string, error) {
	if m.ingredient == nil {
		return "", errors.New("Missing ingredient")
	}

	text, err := m.ingredient.addIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s meat", text), nil
}

type onion struct {
	ingredient ingredientAdd
}

func (o *onion) addIngredient() (string, error) {
	if o.ingredient == nil {
		return "", errors.New("Missing ingredient")
	}

	text, err := o.ingredient.addIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s onions", text), nil
}
