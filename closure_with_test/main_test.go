package main

import "testing"

func TestCreateAddFunc(t *testing.T) {
	inputA := 5
	inputB := 7
	addFive := createAddFunc(inputA)
	expected := 12
	actual := addFive(inputB)
	if actual != expected {
		t.Errorf("createAddFunc(%d)(%d) should return %d, got %d", inputA, inputB, expected, actual)
	}

	inputA = 8
	inputB = 10
	addEight := createAddFunc(inputA)
	expected = 18
	actual = addEight(inputB)
	if actual != expected {
		t.Errorf("createAddFunc(%d)(%d) should return %d, got %d", inputA, inputB, expected, actual)
	}
}
