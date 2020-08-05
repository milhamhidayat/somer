package mariadb

import (
	"database/sql"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
)

type migration struct {
	Migrate *migrate.Migrate
}

// Up is function to run database migrations
func (m *migration) Up() (bool, error) {
	err := m.Migrate.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return true, nil
		}
		return false, err
	}
	return true, nil
}

// Down is function for migrate database to previous version
func (m *migration) Down() (bool, error) {
	err := m.Migrate.Down()
	if err != nil {
		return false, err
	}
	return true, nil
}

func runMigration(dbConn *sql.DB, migrationsFolder string) (*migration, error) {
	dataPath := make([]string, 2)
	dataPath[0] = "file://"
	dataPath[1] = migrationsFolder

	pathToMigrate := strings.Join(dataPath, "")

	driver, err := _mysql.WithInstance(dbConn, &_mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(pathToMigrate, "mysql", driver)
	if err != nil {
		return nil, err
	}
	return &migration{Migrate: m}, nil
}
