package postgresql

import (
	"github.com/geraldie900/todo-app/app/model"
	"github.com/geraldie900/todo-app/config"
	"log"
	"testing"
)

func initPostgre() {
	PostgreDB = NewDatabase(
		config.TodoConfig.PostgreSQLHost,
		config.TodoConfig.PostgreSQLPort,
		config.TodoConfig.PostgreSQLUser,
		config.TodoConfig.PostgreSQLPassword,
		config.TodoConfig.PostgreSQLDBName,
	)
	PostgreDB.AuthDatabase()
	PostgreDB.AutoMigrate()
}

func TestGetWhere(t *testing.T) {
	initPostgre()

	todo := model.Todo{}
	where := "number = ? AND title = ?"
	args := []interface{}{"3", "test 1"}
	PostgreDB.Connection.Where(where, args).Find(&todo)

	log.Println(todo)
}
