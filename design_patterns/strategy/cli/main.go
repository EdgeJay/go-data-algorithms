package main

import (
	"flag"
	"log"
	"os"

	"github.com/edgejay/go-data-algorithms/design_patterns/strategy/shapes"
)

func main() {
	output := flag.String("output", "console", `Possible values are "console" (default) and "image".`)
	flag.Parse()

	activeStrategy, err := shapes.NewPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case shapes.ConsoleStrategy:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.ImageStrategy:
		w, err := os.Create("./image.jpg")
		if err != nil {
			log.Fatalf("Error opening image: %s", err.Error())
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
