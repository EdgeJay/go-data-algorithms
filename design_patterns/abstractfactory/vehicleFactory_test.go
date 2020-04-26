package abstractfactory

import "testing"

func TestGetVehicleFactoryForCar(t *testing.T) {
	factory, err := getVehicleFactory(carFactoryType)
	if err != nil {
		t.Fatalf("Expected error to be nil, got %#v", err)
	}

	outputVehicle, outputErr := factory.newVehicle(familyCarType)
	if outputErr != nil {
		t.Fatalf("Expected error to be nil, got %#v", err)
	}

	fc, ok := outputVehicle.(car)
	if !ok {
		t.Fatal("Expected output vehicle to be casted to car type")
	}

	if fc.getDoors() != 5 {
		t.Errorf("Expected getDoors to return 5, got %d", fc.getDoors())
	}
}

func TestGetVehicleFactoryError(t *testing.T) {
	_, err := getVehicleFactory(10)
	if err == nil {
		t.Errorf("Expected error to be returned, got %#v", err)
	}
}
