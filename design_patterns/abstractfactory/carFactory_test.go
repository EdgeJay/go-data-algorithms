package abstractfactory

import "testing"

func TestCarFactoryReturnFamilyCar(t *testing.T) {
	factory := carFactory{}
	outputCar, err := factory.newVehicle(familyCarType)

	if err != nil {
		t.Fatalf("Expected err to be nil, got %#v", err)
	}

	if outputCar == nil {
		t.Fatal("Expected car to be not nil")
	}

	if outputCar.getSeats() != 7 {
		t.Errorf("Expected getSeats to return 7, got %d", outputCar.getSeats())
	}

	fc, ok := outputCar.(car)
	if !ok {
		t.Fatal("Expected car to be of familyCar type")
	}

	if fc.getDoors() != 5 {
		t.Errorf("Expected getDoors to return 5, got %d", fc.getDoors())
	}
}
