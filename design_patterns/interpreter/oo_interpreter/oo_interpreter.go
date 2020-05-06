package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	operatorSum      = "sum"
	operatorSubtract = "sub"
)

type Interpreter interface {
	Read() int
}

type value int

func (v *value) Read() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (op *operationSum) Read() int {
	return op.Left.Read() + op.Right.Read()
}

type operationSubtract struct {
	Left  Interpreter
	Right Interpreter
}

func (op *operationSubtract) Read() int {
	return op.Left.Read() - op.Right.Read()
}

func operatorFactory(op string, left, right Interpreter) Interpreter {
	switch op {
	case operatorSum:
		return &operationSum{
			Left:  left,
			Right: right,
		}
	case operatorSubtract:
		return &operationSubtract{
			Left:  left,
			Right: right,
		}
	}
	return nil
}

type polishNotationStack []Interpreter

func (p *polishNotationStack) Push(s Interpreter) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() Interpreter {
	count := len(*p)

	if count > 0 {
		temp := (*p)[count-1]
		*p = (*p)[:count-1]
		return temp
	}

	return nil
}

func isOperator(op string) bool {
	return op == operatorSum || op == operatorSubtract
}

func main() {
	stack := polishNotationStack{}
	operators := strings.Split("3 4 sum 2 sub", " ")

	for _, o := range operators {
		if isOperator(o) {
			right := stack.Pop()
			left := stack.Pop()
			operator := operatorFactory(o, left, right)
			res := value(operator.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(o)
			if err != nil {
				panic(err)
			}
			temp := value(val)
			stack.Push(&temp)
		}
	}

	fmt.Printf("Result: %d\n", stack.Pop().Read())
}
