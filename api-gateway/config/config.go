
package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	UserServiceAddr  string `envconfig:"USER_SERVICE_ADDR"`
	TodoServiceAddr  string `envconfig:"TODO_SERVICE_ADDR"`
	RedisServiceAddr string `envconfig:"REDIS_SERVICE_ADDR"`
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}