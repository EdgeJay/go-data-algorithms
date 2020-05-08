package main

import (
	"fmt"
	"strings"
)

func toUpperSync(word string, f func(v string)) {
	f(strings.ToUpper(word))
}

func main() {
	toUpperSync("Hello callbacks!", func(v string) {
		fmt.Printf("From callback: %s\n", v)
	})
}
