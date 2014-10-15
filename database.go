package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"
)

var dbmap *gorp.DbMap

func Connect(conninfo string) {
	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	log.Println("Connected to database")
}

func Cleanup() {
	err := dbmap.Db.Close()

	if err != nil {
		log.Println(err)
	}

	log.Println("Closed database connection")
}
