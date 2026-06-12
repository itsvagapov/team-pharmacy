package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/config"
	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/repository"
	"github.com/itsvagapov/team-pharmacy/internal/service"
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

    err := db.AutoMigrate(&models.User{})
    if err != nil{
        fmt.Println("Миграция не удалась", err)
    }

    userRepo := repository.NewUserRepository(db)

    userService := service.NewUserService(userRepo)
    userHandler := service.NewUserHandler(userSrv)
}