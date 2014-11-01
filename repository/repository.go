package repository

import (
	"errors"

	"github.com/jmoiron/modl"

	"github.com/zacharycarter/monster-codex/entity"
)

var (
	monsterNotFound  = errors.New("monster not found")
	monstersNotFound = errors.New("no monsters found")
)

type Repository interface {
	Get(id uint64) (*interface{}, error)
	List() ([]*interface{}, error)
}

type MonsterRepository struct {
	Executor modl.SqlExecutor
}

func (r *MonsterRepository) Get(id uint64) (*entity.Monster, error) {
	var monster *entity.Monster

	if err := r.Executor.SelectOne(&monster, `SELECT * FROM monster WHERE id=$1;`, id); err != nil {
		return nil, err
	}

	if monster == nil {
		return nil, monsterNotFound
	}

	return monster, nil
}

func (r *MonsterRepository) List() ([]*entity.Monster, error) {
	var monsters []*entity.Monster

	if err := r.Executor.Select(&monsters, `SELECT * FROM monster order by name asc`); err != nil {
		return nil, err
	}

	if len(monsters) == 0 {
		return nil, monstersNotFound
	}

	return monsters, nil
}
