package composite

import "fmt"

type person struct {
	name string
}

type iperson interface {
	getName() string
}

type iswimmer interface {
	swim()
}

type athlete struct {
}

func (a *athlete) train() {
	fmt.Println("I train!")
}

type swimmer struct {
	person
	athlete
}

func (s *swimmer) printName() {
	fmt.Printf("My name is %s\n", s.name)
}

func (s *swimmer) swim() {
	fmt.Println("I swim!")
}

type triathlete struct {
	iperson
	iswimmer
	athlete
	name string
}

func (t *triathlete) getName() string {
	return t.name
}

func (t *triathlete) swim() {
	fmt.Println("I swim!")
}

func (t *triathlete) train() {
	fmt.Println("I train!")
}
