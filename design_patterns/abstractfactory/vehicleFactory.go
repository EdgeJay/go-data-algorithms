package abstractfactory

import "errors"

const (
	carFactoryType = iota + 1
)

type vehicleFactory interface {
	newVehicle(v int) (vehicle, error)
}

func getVehicleFactory(factoryType int) (factory vehicleFactory, err error) {
	switch factoryType {
	case carFactoryType:
		factory = &carFactory{}
	default:
		err = errors.New("Invalid factory type")
	}
	return
}
