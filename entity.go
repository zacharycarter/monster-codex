package main

type Entity interface {
	Equals(a, b []Entity) bool
}

type Monster struct {
	Id          int
	Name        string
	Description string
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
