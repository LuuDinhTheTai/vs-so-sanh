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
	URI            string
	DBName         string
	CollectionName string
}

func LoadEnv() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	uri := os.Getenv("DATABASE_URI")
	dbName := os.Getenv("DATABASE_NAME")
	collectionName := os.Getenv("DATABASE_COLLECTION_NAME")

	return Config{
		Server: ServerConfig{
			Host: host,
			Port: port,
		},
		Database: DatabaseConfig{
			URI:            uri,
			DBName:         dbName,
			CollectionName: collectionName,
		},
	}
}
