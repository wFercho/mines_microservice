package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgres    PostgresConf
	MongoDB     MongoDBConf
	AppPort     string
	MQTTBroker  string
	Environment string
}

func LoadConfig() *Config {

	if err := godotenv.Load(".env.local"); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}

	return &Config{
		Postgres: PostgresConf{
			DBUser:     os.Getenv("POSTGRES_USER"),
			DBPassword: os.Getenv("POSTGRES_PASSWORD"),
			DBName:     os.Getenv("POSTGRES_DB"),
			DBHost:     os.Getenv("POSTGRES_HOST"),
			DBPort:     os.Getenv("POSTGRES_PORT"),
		},
		MongoDB: MongoDBConf{
			DBUser:     os.Getenv("MONGO_USERNAME"),
			DBPassword: os.Getenv("MONGO_PASSWORD"),
			DBName:     os.Getenv("MONGO_DATABASE"),
			DBHost:     os.Getenv("MONGO_HOST"),
			DBPort:     os.Getenv("MONGO_PORT"),
		},
		AppPort:     os.Getenv("APP_PORT"),
		MQTTBroker:  os.Getenv("MQTT_BROKER"),
		Environment: os.Getenv("ENVIRONMENT"),
	}
}
