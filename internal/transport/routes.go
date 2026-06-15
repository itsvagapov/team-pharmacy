package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/repository"
	"github.com/itsvagapov/team-pharmacy/internal/service"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	categoryRepo := repository.NewCategoryRepository(db)

	categoryService := service.NewCategoryService(categoryRepo)

	categoryHandler := NewCategoryHandler(categoryService)

	router.POST("/categories", categoryHandler.CreateCategory)
	router.GET("/categories", categoryHandler.GetAllCategories)

	///

}

func RegisterUserRoutes(router *gin.Engine, userService service.UserService) {

	userHandler := NewUserHandler(userService)

	userHandler.RegisterRoutes(router)
}

///
// userRepo := repository.NewUserRepository(db)
// userService := service.NewUserService(userRepo)
// userHandler := NewUserHandler(userService)

// router.POST("/users", userHandler.Register)
// router.GET("/users/:id", userHandler.GetByID)
