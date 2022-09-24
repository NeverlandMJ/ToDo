
package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

)
// Config is used to get cofigurations of Postgres
type Config struct {
	Host                   string `envconfig:"HOST" required:"true"`
	Port                   string `envconfig:"PORT" required:"true"`
	PostgresHost           string `envconfig:"POSTGRES_HOST" required:"true"`
	PostgresPort           string `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresUser           string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPassword       string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresDB             string `envconfig:"POSTGRES_DB" required:"true"`
	PostgresMigrationsPath string `envconfig:"POSTGRES_MIGRATIONS_PATH" required:"true"`
}

// Load loads configurations values from os
func Load() (Config, error) {
	return Config{
		Host:             os.Getenv("HOST"),
		Port:             os.Getenv("PORT"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresMigrationsPath: os.Getenv("POSTGRES_MIGRATIONS_PATH"),
	}, nil
}