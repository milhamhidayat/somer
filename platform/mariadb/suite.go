package mariadb

import (
	"database/sql"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Suite is struct for mariadb test suite
type Suite struct {
	suite.Suite
	DBConn          *sql.DB
	DSN             string
	DBName          string
	Migration       *migration
	MigrationFolder string
}

// New will initialize test suite
func (s *Suite) New() {
	var err error

	s.DBConn, err = sql.Open("mysql", s.DSN)
	require.NoError(s.T(), err)

	err = s.DBConn.Ping()
	require.NoError(s.T(), err)

	s.Migration, err = runMigration(s.DBConn, s.MigrationFolder)
	require.NoError(s.T(), err)
}

// TearDown will close db connection
func (s *Suite) TearDown() {
	s.DBConn.Close()
}
