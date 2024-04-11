package main

import (
	"api_gateway_svc/config"
	clientsetup "api_gateway_svc/stream/client_setup"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("err loading config", err)
	}

	r := gin.Default()

	clientsetup.RegisterRoutes(r, &config)

	r.Run()
}
