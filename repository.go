package main

import (
	"errors"
	"github.com/coopernurse/gorp"
)

type EntityFinder interface {
}

type Repository interface {
	EntityFinder
	All() ([]Monster, error)
}

type MonsterRepository struct {
	*gorp.DbMap
}

var NotFoundError = errors.New("Entity not found")
var DuplicateNameError = errors.New("This name is already in use")

func NewMonsterRepository() Repository {
	return &MonsterRepository{dbmap}
}

func (r MonsterRepository) All() ([]Monster, error) {
	var monsters []Monster
	_, err := dbmap.Select(&monsters, "select * from monster order by name")
	return monsters, err
}
