package db

import (
	"sync"

	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/zacharycarter/monster-codex/config"
	"github.com/zacharycarter/monster-codex/util"
)

var (
	DbMap   = &modl.DbMap{Dialect: modl.PostgresDialect{}}
	connect sync.Once
)

func Connect() {
	connect.Do(func() {
		var err error
		DbMap.Dbx, err = sqlx.Open("postgres", config.ConnectionString)
		util.CheckErr(err, "Error connecting to database: ")
		DbMap.Db = DbMap.Dbx.DB
	})
}

func Initialize() {
	err := DbMap.CreateTablesIfNotExists()
	util.CheckErr(err, "Error creating tables: ")
}

func Close() {
	err := DbMap.Db.Close()
	util.CheckErr(err, "Error closing database: ")
}
