package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)

func RegisterRoutes(router *gin.Engine, categoryService service.CategoryService, subcategoryService service.SubcategoryService, userService service.UserService) { // userService service.User....
	categoryHandler := NewCategoryHandler(categoryService)
	subcategoryHandler := NewSubcategoryHandler(subcategoryService)
	userHandler := NewUserHandler(userService)

	categoryHandler.RegisterRoutes(router)
	subcategoryHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
}
