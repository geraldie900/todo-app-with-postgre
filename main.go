package main

import (
	"github.com/geraldie900/todo-app/app/handler/gofiber"
	"github.com/geraldie900/todo-app/config"
	"github.com/geraldie900/todo-app/config/database/postgresql"
	"github.com/geraldie900/todo-app/config/logger"
)

func main() {
	logger.InitLogging()

	postgresql.PostgreDB = postgresql.NewDatabase(
		config.TodoConfig.PostgreSQLHost,
		config.TodoConfig.PostgreSQLPort,
		config.TodoConfig.PostgreSQLUser,
		config.TodoConfig.PostgreSQLPassword,
		config.TodoConfig.PostgreSQLDBName,
	)
	postgresql.PostgreDB.AuthDatabase()
	postgresql.PostgreDB.AutoMigrate()

	gofiber.InitGofiber()
}
