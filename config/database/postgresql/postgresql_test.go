package postgresql

import (
	"github.com/geraldie900/todo-app/config"
	"testing"
)

func TestConnectToPostgre(t *testing.T) {
	postgreDB := NewDatabase(
		config.TodoConfig.PostgreSQLHost,
		config.TodoConfig.PostgreSQLPort,
		config.TodoConfig.PostgreSQLUser,
		config.TodoConfig.PostgreSQLPassword,
		config.TodoConfig.PostgreSQLDBName,
	)
	postgreDB.AuthDatabase()
}

func TestAutoMigrator(t *testing.T) {
	postgreDB := NewDatabase(
		config.TodoConfig.PostgreSQLHost,
		config.TodoConfig.PostgreSQLPort,
		config.TodoConfig.PostgreSQLUser,
		config.TodoConfig.PostgreSQLPassword,
		config.TodoConfig.PostgreSQLDBName,
	)
	postgreDB.AuthDatabase()
	postgreDB.AutoMigrate()
}
