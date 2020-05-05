package memento

import "fmt"

type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(index int) (memento, error) {
	if index < 0 || index >= len(c.mementoList) {
		return memento{}, fmt.Errorf("Invalid index: %d", index)
	}

	return c.mementoList[index], nil
}
