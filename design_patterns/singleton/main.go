package singleton

import "sync"

type counterSingleton struct {
	value int
}

func (counter *counterSingleton) AddOne() int {
	counter.value++
	return counter.value
}

func (counter *counterSingleton) Value() int {
	return counter.value
}

var instance *counterSingleton
var once sync.Once

// GetInstance returns singleton instance of counter
func GetInstance() *counterSingleton {
	once.Do(func() {
		instance = &counterSingleton{}
	})
	return instance
}
