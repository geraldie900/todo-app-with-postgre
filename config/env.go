package config

import (
	"github.com/geraldie900/todo-app/config/utils"
	structvalidator "github.com/geraldie900/todo-app/config/validator"
	"log"

	"github.com/spf13/viper"
)

type EnvConfigApp struct {
	AppPort  string `mapstructure:"APP_PORT" validate:"required"`
	LogLevel string `mapstructure:"LOG_LEVEL" validate:"required"`
}

type EnvConfigTodoApp struct {
	PostgreSQLHost     string `mapstructure:"POSTGRE_SQL_HOST" validate:"required"`
	PostgreSQLPort     string `mapstructure:"POSTGRE_SQL_PORT" validate:"required"`
	PostgreSQLUser     string `mapstructure:"POSTGRE_SQL_USER" validate:"required"`
	PostgreSQLPassword string `mapstructure:"POSTGRE_SQL_PASSWORD" validate:"required"`
	PostgreSQLDBName   string `mapstructure:"POSTGRE_SQL_DB_NAME" validate:"required"`
}

var (
	AppConfig  EnvConfigApp
	TodoConfig EnvConfigTodoApp
)

func init() {
	// path for local test
	viper.AddConfigPath("../../../.")
	// path for local run
	viper.AddConfigPath("../.")
	// path for server run
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("env")

	// replace value from config.env if from OS available
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("[init] read config file failed:", err)
	}

	// marshal to env config app
	if err = viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalln("[init] viper unmarshal env config app", err)
	}
	if err = structvalidator.StructValidator(AppConfig); err != nil {
		log.Fatalln("[init] missing env config value app", err)
	}
	utils.StructIterator(AppConfig)

	// marshal to env to-do app
	if err = viper.Unmarshal(&TodoConfig); err != nil {
		log.Fatalln("[init] viper unmarshal env config app", err)
	}
	if err = structvalidator.StructValidator(TodoConfig); err != nil {
		log.Fatalln("[init] missing env config value app", err)
	}
	utils.StructIterator(TodoConfig)
}
