package postgresql

import (
	"fmt"
	"github.com/geraldie900/todo-app/app/model"
	"github.com/geraldie900/todo-app/config/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBConnector interface {
	AuthDatabase()
	AutoMigrate()
}

type DatabaseConfig struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	Connection       *gorm.DB
}

var (
	PostgreDB *DatabaseConfig
)

func NewDatabase(host, port, username, password, dbName string) *DatabaseConfig {
	return &DatabaseConfig{
		DatabaseHost:     host,
		DatabasePort:     port,
		DatabaseUsername: username,
		DatabasePassword: password,
		DatabaseName:     dbName,
	}
}

func (dbc *DatabaseConfig) AuthDatabase() {
	logging := logger.Logger{
		FunctionName: "AuthDatabase",
	}
	logging.LogInfo("init postgre sql")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbc.DatabaseHost,
		dbc.DatabaseUsername,
		dbc.DatabasePassword,
		dbc.DatabaseName,
		dbc.DatabasePort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.LogError("gorm open postgre", err)
		log.Fatal()
	}

	dbc.Connection = db
	log.Println(dbc.Connection)
	logging.LogInfo("connected to postgre sql")
}

func (dbc *DatabaseConfig) AutoMigrate() {
	logging := logger.Logger{
		FunctionName: "AutoMigrate",
	}
	logging.LogInfo("running auto migrate postgre")

	err := dbc.Connection.AutoMigrate(&model.Todo{})
	if err != nil {
		logging.LogError("auto migrate fail", err)
		log.Fatal()
	}

	logging.LogInfo("auto migrate finish")
}
