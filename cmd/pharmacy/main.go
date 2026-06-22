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

	if err := db.AutoMigrate(&models.Category{}, &models.Subcategory{}, &models.Review{}, &models.Medicine{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	categoryRepo := repository.NewCategoryRepository(db)
	subcategoryRepo := repository.NewSubcategoryRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	medicineRepo := repository.NewMedicineRepository(db)

	categoryService := service.NewCategoryService(categoryRepo)
	subcategoryService := service.NewSubcategoryService(subcategoryRepo, categoryRepo)
	reviewService := service.NewReviewService(reviewRepo, medicineRepo)
	medicineService := service.NewMedicineService(medicineRepo, categoryRepo, subcategoryRepo)

	router := gin.Default()

	transport.RegisterRoutes(router, categoryService, subcategoryService, reviewService, medicineService)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
