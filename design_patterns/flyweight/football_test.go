package flyweight

import (
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := newTeamFactory()

	teamA1 := factory.getTeam(teamA)
	if teamA1 == nil {
		t.Error("Expected team to be not nil")
	}

	teamA2 := factory.getTeam(teamA)
	if teamA2 == nil {
		t.Error("Expected team to be not nil")
	}

	if teamA1 != teamA2 {
		t.Error("Expected teamA objects to be the same")
	}

	if factory.getNumberOfObjects() != 1 {
		t.Errorf("Expected number of objects in factory to be 1, got %d", factory.getNumberOfObjects())
	}
}

func Test_HighVolume(t *testing.T) {
	factory := newTeamFactory()
	teams := make([]*team, 500000*2)

	for idx := 0; idx < 500000; idx++ {
		teams[idx] = factory.getTeam(teamA)
	}

	for idx := 500000; idx < 500000*2; idx++ {
		teams[idx] = factory.getTeam(teamB)
	}

	if factory.getNumberOfObjects() != 2 {
		t.Errorf("Expected number of objects to be 2, got %d", factory.getNumberOfObjects())
	}

	for idx := 0; idx < 3; idx++ {
		t.Logf("Pointer %d points to %p and is located in %p", idx, teams[idx], &teams[idx])
	}
}
