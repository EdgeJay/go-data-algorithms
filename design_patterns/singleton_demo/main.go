package main

import (
	"fmt"

	"github.com/edgejay/go-data-algorithms/design_patterns/singleton"
)

func main() {
	singletonCounter := singleton.GetInstance()
	fmt.Println("singletonCounter.AddOne() : ", singletonCounter.AddOne())
	fmt.Println("singletonCounter.AddOne() : ", singletonCounter.AddOne())
	fmt.Println("singletonCounter.AddOne() : ", singletonCounter.AddOne())

	anotherCounter := singleton.GetInstance()
	fmt.Println("anotherCounter.Value() : ", anotherCounter.Value())
}
