package configuration

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	URI    string
	DBName string
}

func LoadEnv() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//appEnv := os.Getenv("APP_ENV")

	//ginMode := os.Getenv("GIN_MODE")

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	//dbUser := os.Getenv("DATABASE_USER")
	//dbPassword := os.Getenv("DATABASE_PASSWORD")
	uri := os.Getenv("DATABASE_URI")
	dbName := os.Getenv("DATABASE_NAME")

	return Config{
		Server: ServerConfig{
			Host: host,
			Port: port,
		},
		Database: DatabaseConfig{
			URI:    uri,
			DBName: dbName,
		},
	}
}
