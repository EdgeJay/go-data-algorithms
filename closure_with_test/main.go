package main

import "fmt"

func createAddFunc(numA int) func(int) int {
	return func(numB int) int {
		return numA + numB
	}
}

func main() {
	addFive := createAddFunc(5)
	fmt.Printf("addFive (5 + 7 = %d)\n", addFive(7))
}
