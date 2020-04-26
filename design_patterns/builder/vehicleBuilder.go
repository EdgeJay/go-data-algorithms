package builder

type vehicle struct {
	structure string
	wheels    int
	seats     int
}

type builder interface {
	setStructure() builder
	setWheels() builder
	setSeats() builder
	getVehicle() vehicle
}

type carBuilder struct {
	veh vehicle
}

func (builder *carBuilder) setStructure() builder {
	builder.veh.structure = "car"
	return builder
}

func (builder *carBuilder) setWheels() builder {
	builder.veh.wheels = 4
	return builder
}

func (builder *carBuilder) setSeats() builder {
	builder.veh.seats = 5
	return builder
}

func (builder *carBuilder) getVehicle() vehicle {
	return builder.veh
}

type manufacturingComplex struct {
	builder builder
}

func (m *manufacturingComplex) construct() vehicle {
	return m.builder.setStructure().setWheels().setSeats().getVehicle()
}

func (m *manufacturingComplex) setBuilder(b builder) {
	m.builder = b
}
