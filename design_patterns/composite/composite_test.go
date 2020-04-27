package composite

import (
	"fmt"
	"testing"
)

func TestSwimmer(t *testing.T) {
	aSwimmer := swimmer{
		person: person{"John Smith"},
	}
	aSwimmer.printName()
	aSwimmer.train()
	aSwimmer.swim()
	t.Log("Swimmer test completed")
}

func TestTriathlete(t *testing.T) {
	aTriathlete := triathlete{name: "Jane Doe"}
	fmt.Printf("My name is %s\n", aTriathlete.getName())
	aTriathlete.train()
	aTriathlete.swim()
	t.Log("Swimmer test completed")
}
