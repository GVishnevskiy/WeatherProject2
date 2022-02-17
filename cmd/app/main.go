package main

import (
	"../../internal/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// REDIS
// Docker-compose

func main() {
	_ = godotenv.Load()

	router := gin.Default()

	server.StartServer(router)
}
