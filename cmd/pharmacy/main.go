package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/config"
	"github.com/itsvagapov/team-pharmacy/internal/transport"
)

func main() {
	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	transport.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
