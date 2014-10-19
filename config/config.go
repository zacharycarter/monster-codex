package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/zacharycarter/monster-codex/util"
)

type ConnectionStrings struct {
	Development string
	Test        string
	Production  string
}

var (
	environment      = getEnvironment()
	ConnectionString = getConnectionString()
	Port             = getPort()
)

func getEnvironment() string {
	env := os.Getenv("ENV")

	if env == "" {
		return "development"
	} else {
		return env
	}
}

func getConnectionString() string {
	database_url := os.Getenv("DATABASE_URL")

	if database_url != "" {
		return database_url
	} else {
		connectionStrings := &ConnectionStrings{}

		environmentData, err := ioutil.ReadFile("./config/database.yaml")

		util.CheckErr(err, "Error reading database.yaml: ")

		err = yaml.Unmarshal(environmentData, connectionStrings)

		util.CheckErr(err, "Error unmarshalling connection strings: ")

		switch environment {
		case "development":
			database_url = connectionStrings.Development
		case "test":
			database_url = connectionStrings.Test
		case "production":
			database_url = connectionStrings.Production
		default:
			log.Fatal("Unrecognized environment: ", environment)
		}
	}
	return database_url
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "9898"
	} else {
		return port
	}
}
