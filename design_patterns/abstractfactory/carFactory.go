package abstractfactory

const (
	familyCarType = iota + 1
	sportsCarType
)

type carFactory struct{}

func (cf *carFactory) newVehicle(v int) (veh vehicle, err error) {
	switch v {
	case familyCarType:
		veh = &familyCar{}
	}
	return
}
