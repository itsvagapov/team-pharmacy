package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)

type SubcategoryHandler struct {
	service service.SubcategoryService
}

func NewSubcategoryHandler(service service.SubcategoryService) *SubcategoryHandler {
	return &SubcategoryHandler{service: service}
}

func (h *SubcategoryHandler) RegisterRoutes(r *gin.Engine) {
	subcategories := r.Group("/categories")
	{
		subcategories.GET("/:id/subcategories", h.GetSubcategoriesByCategoryID)
		subcategories.POST("/:id/subcategories", h.CreateSubcategory)
	}
}

func (h *SubcategoryHandler) CreateSubcategory(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректный id категории",
		})
		return
	}

	var req models.SubcategoryCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректное тело запроса",
		})
		return
	}

	subcategory, err := h.service.CreateSubcategory(
		uint(id),
		req,
	)

	if err != nil {
		if errors.Is(err, service.ErrSubcategoryNameRequired) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, service.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "внутренняя ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusCreated, subcategory)
}

func (h *SubcategoryHandler) GetSubcategoriesByCategoryID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректный id категории",
		})
		return
	}

	subcategories, err := h.service.GetSubcategoriesByCategoryID(
		uint(id),
	)

	if err != nil {
		if errors.Is(err, service.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "внутренняя ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, subcategories)
}
