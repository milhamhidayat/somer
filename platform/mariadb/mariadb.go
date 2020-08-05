package mariadb

import (
	"database/sql"

	"github.com/rs/zerolog/log"
)

// Database where the database connection is saved
type Database struct {
	Primary *sql.DB
}

// Connections is the interface that wraps the basic Open method for mariaDB
type Connections interface {
	Open() *Database
}

type connectionString struct {
	primary string
	app     string
}

// New is initialize mariadb using db connecting string
// Never set it to empty, because it will stop the entire app
func New(primary, app string) Connections {
	return &connectionString{
		primary: primary,
		app:     app,
	}
}

// Open is creating database connection
func (cs *connectionString) Open() *Database {
	log.Info().Fields(map[string]interface{}{
		"platform": "mariadb",
		"app":      cs.app,
	}).Msg("Connecting to mariaDB")

	dbPrimary, err := sql.Open("mysql", cs.primary)
	if err != nil {
		log.Fatal().Fields(map[string]interface{}{
			"connection": cs.primary,
			"err":        err,
		}).Msg("connect to mariaDB")
	}

	err = dbPrimary.Ping()
	if err != nil {
		log.Fatal().Fields(map[string]interface{}{
			"connection": cs.primary,
			"err":        err,
		}).Msg("ping db")
	}

	return &Database{
		Primary: dbPrimary,
	}
}
