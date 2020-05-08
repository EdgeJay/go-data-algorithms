package singleton

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

type singleton struct{}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	ch := make(chan int)
	defer close(ch)

	getCountCh <- ch
	return <-ch
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func init() {
	var count int

	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				break
			}
		}
	}(addCh, getCountCh, quitCh)
}
