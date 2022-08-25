package main

import (
	"github.com/NeverlandMJ/ToDo/api-gateway/service"
	"github.com/NeverlandMJ/ToDo/api-gateway/api"
)

const (
	userServiceURL = "localhost:9000"
	todoServiceURL = "localhost:9001"
)

func main() {
	h := api.NewHandler(service.NewProvider(userServiceURL, todoServiceURL))
	router := api.NewRouter(h)

	router.Run()
}