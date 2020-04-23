package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	counter2 := GetInstance()

	if counter1 != counter2 {
		t.Error("Expect GetInstance to return same counter instance")
	}

	result := counter1.AddOne()
	if result != counter2.Value() {
		t.Error("Expect counter instances to have same values")
	}
}
