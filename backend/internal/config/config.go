package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string

	MysqlUser     string
	MysqlPassword string
	MysqlHost     string
	MysqlPort     string
	MysqlDatabase string
}

func Load() Config {
	godotenv.Load()
	return Config{
		AppName:   "Curly",
		Port:      "8080",
		MysqlUser: os.Getenv("MYSQL_USER"),

		MysqlPassword: os.Getenv("MYSQL_PASSWORD"),

		MysqlHost: os.Getenv("MYSQL_HOST"),

		MysqlPort: os.Getenv("MYSQL_PORT"),

		MysqlDatabase: os.Getenv("MYSQL_DATABASE"),
	}
}
