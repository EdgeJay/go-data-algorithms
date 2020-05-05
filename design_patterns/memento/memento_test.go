package memento

import "testing"

func TestCareTaker_Add(t *testing.T) {
	expectedState := "Idle"
	originator := originator{}
	originator.state = State{Description: expectedState}
	careTaker := careTaker{}
	mem := originator.NewMemento()
	if mem.state.Description != expectedState {
		t.Errorf("Expected %s state, got %s", expectedState, mem.state.Description)
	}

	currentLen := len(careTaker.mementoList)
	careTaker.Add(mem)
	if len(careTaker.mementoList) != currentLen+1 {
		t.Errorf("Expected mementoList to grow to %d, got %d", currentLen+1, len(careTaker.mementoList))
	}
}

func TestCareTaker_Memento(t *testing.T) {
	expectedState := "Idle"
	originator := originator{}
	careTaker := careTaker{}
	originator.state = State{"Idle"}
	careTaker.Add(originator.NewMemento())

	mem, err := careTaker.Memento(0)
	if err != nil {
		t.Fatalf("Expected error to be nil, got %s", err.Error())
	}

	if mem.state.Description != "Idle" {
		t.Errorf("Expected %s state, got %s", expectedState, mem.state.Description)
	}

	mem, err = careTaker.Memento(-1)
	if err == nil {
		t.Fatal("Expected error to be not nil")
	}
}

func TestOriginator_ExtractAndStoreState(t *testing.T) {
	expectedState := "Idle"
	originator := originator{state: State{expectedState}}
	idleMemento := originator.NewMemento()
	originator.state = State{"Working"}
	originator.ExtractAndStoreState(idleMemento)

	if originator.state.Description != expectedState {
		t.Errorf("Expected %s state, got %s", expectedState, originator.state.Description)
	}
}
