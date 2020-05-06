package interpreter

import (
	"strconv"
	"strings"
)

const (
	operatorSum      = "sum"
	operatorSubtract = "sub"
	operatorMultiply = "mul"
	operatorDivide   = "div"
)

type polishNotationStack []int

func (p *polishNotationStack) Push(s int) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() int {
	count := len(*p)

	if count > 0 {
		temp := (*p)[count-1]
		*p = (*p)[:count-1]
		return temp
	}

	return 0
}

func isOperator(op string) bool {
	return op == operatorSum || op == operatorSubtract || op == operatorMultiply || op == operatorDivide
}

func getOperationFunc(op string) func(int, int) int {
	switch op {
	case operatorSum:
		return func(left, right int) int {
			return left + right
		}
	case operatorSubtract:
		return func(left, right int) int {
			return left - right
		}
	case operatorMultiply:
		return func(left, right int) int {
			return left * right
		}
	case operatorDivide:
		return func(left, right int) int {
			return left / right
		}
	default:
		return func(left, right int) int {
			return 0
		}
	}
}

func Calculate(op string) (int, error) {
	stack := polishNotationStack{}
	operators := strings.Split(op, " ")

	for _, o := range operators {
		if isOperator(o) {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := getOperationFunc(o)
			res := mathFunc(left, right)
			stack.Push(res)
		} else {
			val, err := strconv.Atoi(o)
			if err != nil {
				return 0, err
			}
			stack.Push(val)
		}
	}

	return stack.Pop(), nil
}
