package abstractfactory

type sportsCar struct{}

func (car *sportsCar) getDoors() int {
	return 2
}

func (car *sportsCar) getWheels() int {
	return 4
}

func (car *sportsCar) getSeats() int {
	return 2
}
