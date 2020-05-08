package observer

import (
	"fmt"
	"testing"
)

type testObserver struct {
	ID      int
	Message string
}

func (p *testObserver) Notify(m string) {
	fmt.Printf(`Observer %d: Message "%s" received\n`, p.ID, p.Message)
	p.Message = m
}

func TestSubject(t *testing.T) {
	defaultMessage := "default"
	testObserver1 := &testObserver{1, defaultMessage}
	testObserver2 := &testObserver{2, defaultMessage}
	testObserver3 := &testObserver{3, defaultMessage}
	publisher := Publisher{}

	t.Run("AddObserver", func(t *testing.T) {
		publisher.AddObserver(testObserver1)
		publisher.AddObserver(testObserver2)
		publisher.AddObserver(testObserver3)

		if len(publisher.ObserverList) != 3 {
			t.Errorf("Expected length of observers list to be 3, got %d", len(publisher.ObserverList))
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher.RemoveObserver(testObserver2)

		if len(publisher.ObserverList) != 2 {
			t.Errorf("Expected length of observers list to be 2, got %d", len(publisher.ObserverList))
		}

		for _, observer := range publisher.ObserverList {
			testObserver, ok := observer.(*testObserver)
			if !ok {
				t.Error("Observer is not of type *testObserver")
			}

			if testObserver.ID == 2 {
				t.Error("Expected observer with ID 2 to be not present")
			}
		}
	})

	t.Run("NotifyObservers", func(t *testing.T) {
		expectedMessage := "hello world"

		if len(publisher.ObserverList) == 0 {
			t.Fatal("Observer list is empty, nothing to test")
		}

		for _, observer := range publisher.ObserverList {
			testObserver, ok := observer.(*testObserver)
			if !ok {
				t.Error("Observer is not of type *testObserver")
			}

			if testObserver.Message != defaultMessage {
				t.Errorf("Expected observer's message to be of default value, got %s", testObserver.Message)
			}
		}

		publisher.NotifyObservers(expectedMessage)

		for _, observer := range publisher.ObserverList {
			testObserver, _ := observer.(*testObserver)
			if testObserver.Message != expectedMessage {
				t.Errorf("Expected observer's message to be of default value, got %s", testObserver.Message)
			}
		}
	})
}
