package main

import (
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

	if err := db.AutoMigrate(&models.Category{}, &models.Subcategory{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}
	// -TODO: перекинуть юзера
	// err := db.AutoMigrate(&models.User{})
	// if err != nil {
	// 	fmt.Println("Миграция не удалась", err)
	// }

	categoryRepo := repository.NewCategoryRepository(db)
	subcategoryRepo := repository.NewSubcategoryRepository(db)

	userRepo := repository.NewUserRepository(db)
	// userRepo ...

	categoryService := service.NewCategoryService(categoryRepo)
	subcategoryService := service.NewSubcategoryService(subcategoryRepo, categoryRepo)
	userService := service.NewUserService(userRepo)
	// userService ,,,

	router := gin.Default()

	transport.RegisterRoutes(router, categoryService, subcategoryService, userService) // ...userService

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}

	router.Run(":8888")

}
