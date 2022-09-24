package main

import (
	"fmt"
	"log"

	"github.com/NeverlandMJ/ToDo/api-gateway/api"
	"github.com/NeverlandMJ/ToDo/api-gateway/config"
	"github.com/NeverlandMJ/ToDo/api-gateway/service"
)

// default hosts urls for services
const (
	userServiceURL = "localhost:9000"
	todoServiceURL = "localhost:9001"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Panicln("failed to load config", err)
	}

	fmt.Println(cfg)

	h := api.NewHandler(service.NewProvider(cfg.UserServiceAddr, cfg.TodoServiceAddr, cfg.RedisServiceAddr))

	router := api.NewRouter(h)

	router.Run()
}
