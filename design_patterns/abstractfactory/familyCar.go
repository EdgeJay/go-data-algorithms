package abstractfactory

type familyCar struct{}

func (car *familyCar) getDoors() int {
	return 5
}

func (car *familyCar) getWheels() int {
	return 4
}

func (car *familyCar) getSeats() int {
	return 7
}
