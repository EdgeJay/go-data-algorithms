package interpreter

import "testing"

func TestCalculate(t *testing.T) {
	expectedResult := 4
	tempOperation := "5 3 sub 8 mul 4 sum 5 div"
	res, err := Calculate(tempOperation)

	if err != nil {
		t.Errorf("Expected error to be nil, got %s", err.Error())
	}

	if res != expectedResult {
		t.Errorf("Expected result to be %d, got %d", expectedResult, res)
	}

	expectedResult = 5
	tempOperation = "3 4 sum 2 sub"
	res, err = Calculate(tempOperation)

	if err != nil {
		t.Errorf("Expected error to be nil, got %s", err.Error())
	}

	if res != expectedResult {
		t.Errorf("Expected result to be %d, got %d", expectedResult, res)
	}
}
