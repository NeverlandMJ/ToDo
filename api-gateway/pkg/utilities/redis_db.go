package utilities

import (
	"github.com/go-redis/redis"
)

// NewRedisClient creates redis client
func NewRedisClient(addr string) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
