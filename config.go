package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type DatabaseEnvironments struct {
	Development string
	Test        string
	Production  string
}

var connectionString string

func DatabaseConnectionString() string {
	if connectionString == "" {
		initConnectionString()
	}

	return connectionString
}

func Port() string {
	port := os.Getenv("PORT")

	if port != "" {
		return port
	}

	return "8080"
}

// Uses connection info from $DATABASE_URL if
// available.
// If not, use info from PROJECT_ROOT/db/database.yaml
func initConnectionString() {
	database_url := os.Getenv("DATABASE_URL")

	if database_url != "" {
		connectionString = database_url
	} else {
		databaseEnvironments := new(DatabaseEnvironments)

		environmentData, err := ioutil.ReadFile("./database.yaml")

		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(environmentData, &databaseEnvironments)

		if err != nil {
			panic(err)
		}

		env := Env().String()

		switch env {
		case "development":
			connectionString = databaseEnvironments.Development
		case "test":
			connectionString = databaseEnvironments.Test
		case "production":
			connectionString = databaseEnvironments.Production
		default:
			panic(fmt.Sprintf("No database config for environment %s", env))
		}
	}
}
