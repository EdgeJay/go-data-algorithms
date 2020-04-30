package flyweight

import "time"

const (
	teamA = iota
	teamB
)

type player struct {
	name         string
	surname      string
	previousTeam uint64
	photo        []byte
}

type match struct {
	date          time.Time
	visitorID     uint64
	localID       uint64
	localScore    byte
	visitorScore  byte
	localShoots   uint16
	visitorShoots uint16
}

type historicalData struct {
	year          uint8
	leagueResults []match
}

type team struct {
	id             uint64
	name           string
	shield         []byte
	players        []player
	historicalData []historicalData
}

type teamFlyweightFactory struct {
	createdTeams map[int]*team
}

func newTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*team),
	}
}

func getTeamFactory(id int) team {
	switch id {
	case teamB:
		return team{
			id:   2,
			name: "TEAM_B",
		}
	default:
		return team{
			id:   1,
			name: "TEAM_A",
		}
	}
}

func (t *teamFlyweightFactory) getTeam(id int) *team {
	if t.createdTeams[id] != nil {
		return t.createdTeams[id]
	}

	team := getTeamFactory(id)
	t.createdTeams[id] = &team

	return t.createdTeams[id]
}

func (t *teamFlyweightFactory) getNumberOfObjects() int {
	return len(t.createdTeams)
}
