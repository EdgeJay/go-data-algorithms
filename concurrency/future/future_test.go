package future

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func timeout(t *testing.T, waitGroup *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("timeout")
	t.Fail()
	waitGroup.Done()
}

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}

	t.Run("Success result", func(t *testing.T) {
		var waitGroup sync.WaitGroup
		waitGroup.Add(1)

		future.Success(func(s string) {
			t.Log(s)
			waitGroup.Done()
		}).Fail(func(e error) {
			t.Error("Expected future to be successful")
			waitGroup.Done()
		})

		future.Execute(func() (string, error) {
			return "Hello world!", nil
		})

		waitGroup.Wait()
	})

	t.Run("Failed result", func(t *testing.T) {
		var waitGroup sync.WaitGroup
		waitGroup.Add(1)

		future.Success(func(s string) {
			t.Error("Expected future to fail")
			waitGroup.Done()
		}).Fail(func(e error) {
			t.Log(e)
			waitGroup.Done()
		})

		future.Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})

		waitGroup.Wait()
	})

	t.Run("Closure success result", func(t *testing.T) {
		var waitGroup sync.WaitGroup
		waitGroup.Add(1)

		go timeout(t, &waitGroup)

		future.Success(func(s string) {
			t.Log(s)
			waitGroup.Done()
		}).Fail(func(e error) {
			t.Error("Expected future to be successful")
			waitGroup.Done()
		})

		future.Execute(setContext("Hello!"))

		waitGroup.Wait()
	})
}
