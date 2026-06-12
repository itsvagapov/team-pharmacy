package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)

func RegisterRoutes(router *gin.Engine, categoryService service.CategoryService, subcategoryService service.SubcategoryService) {
	categoryHandler := NewCategoryHandler(categoryService)
	subcategoryHandler := NewSubcategoryHandler(subcategoryService)

	categoryHandler.RegisterRoutes(router)
	subcategoryHandler.RegisterRoutes(router)
}
