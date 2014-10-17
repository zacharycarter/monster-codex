package main

type Entity interface {
	Equals(a, b []Entity) bool
}

type Monster struct {
	Id           int
	Name         string
	Keywords     []uint8
	Description  string
	Hint         string
	Power        int
	Charisma     int
	Intelligence int
	Finesse      int
	Talents      []uint8
	Strikes      int
}

func (m Monster) Equals(a, b []Entity) bool {
	if len(a) != len(b) {
		return false
	}

	for i, monster := range a {
		if b[i] != monster {
			return false
		}
	}

	return true
}
