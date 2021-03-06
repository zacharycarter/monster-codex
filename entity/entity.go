package entity

import (
	"github.com/zacharycarter/monster-codex/db"
)

type Monster struct {
	Id           uint64
	Name         string
	Keywords     db.StringSlice
	Description  string
	Hint         string
	Power        int8
	Charisma     int8
	Intelligence int8
	Finesse      int8
	Talents      db.StringSlice
	Strikes      int8
	Image        string
}

func init() {
	db.DbMap.AddTableWithName(Monster{}, "monster").SetKeys(true, "id")
}
