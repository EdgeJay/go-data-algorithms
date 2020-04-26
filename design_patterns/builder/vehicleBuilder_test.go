package builder

import "testing"

func TestConstruct(t *testing.T) {
	factory := manufacturingComplex{}
	factory.setBuilder(&carBuilder{})
	vehicle := factory.construct()

	if vehicle.structure != "car" {
		t.Errorf("Expected structure of vehicle constructed using carBuilder to be \"car\", got %s", vehicle.structure)
	}

	if vehicle.wheels != 4 {
		t.Errorf("Expected wheels of vehicle constructed using carBuilder to be 4, got %d", vehicle.wheels)
	}

	if vehicle.seats != 5 {
		t.Errorf("Expected seats of vehicle constructed using carBuilder to be 5, got %d", vehicle.seats)
	}
}
