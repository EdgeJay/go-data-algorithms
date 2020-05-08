package observer

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserverList []Observer
}

func (p *Publisher) AddObserver(ob Observer) {
	p.ObserverList = append(p.ObserverList, ob)
}

func (p *Publisher) RemoveObserver(ob Observer) {
	newList := []Observer{}
	for _, observer := range p.ObserverList {
		if observer != ob {
			newList = append(newList, observer)
		}
	}

	p.ObserverList = newList
}

func (p *Publisher) NotifyObservers(m string) {
	for _, observer := range p.ObserverList {
		observer.Notify(m)
	}
}
