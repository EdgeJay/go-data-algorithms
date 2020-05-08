package singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstances(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	for index := 0; index < n; index++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(time.Millisecond * 10)
	}

	singleton.Stop()
}
