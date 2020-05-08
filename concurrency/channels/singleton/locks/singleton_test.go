package singleton

import (
	"testing"
	"time"
)

func TestStartInstances(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	t.Logf("Before loop, current count is %d\n", singleton.GetCount())

	for index := 0; index < n; index++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	t.Logf("After loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(time.Millisecond * 10)
	}
}
