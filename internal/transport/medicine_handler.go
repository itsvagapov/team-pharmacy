package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)

type MedicineHandler struct {
	service service.MedicineService
}

func NewMedicineHandler(service service.MedicineService) *MedicineHandler {
	return &MedicineHandler{
		service: service,
	}
}

func (h *MedicineHandler) RegisterRoutes(r *gin.Engine) {
	medicines := r.Group("/medicines")
	{
		medicines.GET("", h.GetAllMedicines)
		medicines.GET("/:id", h.GetMedicineByID)
		medicines.POST("", h.CreateMedicine)
		medicines.PATCH("/:id", h.UpdateMedicine)
		medicines.DELETE("/:id", h.DeleteMedicine)
	}
}

func (h *MedicineHandler) CreateMedicine(c *gin.Context) {
	var req models.MedicineCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	medicine, err := h.service.CreateMedicine(req)
	if err != nil {
		if errors.Is(err, service.ErrMedicineNameRequired) ||
			errors.Is(err, service.ErrMedicineDescriptionRequired) ||
			errors.Is(err, service.ErrManufacturerRequired) ||
			errors.Is(err, service.ErrPriceMustBePositive) ||
			errors.Is(err, service.ErrStockQuantityNegative) ||
			errors.Is(err, service.ErrCategoryNotFound) ||
			errors.Is(err, service.ErrSubcategoryNotFound) {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, medicine)
}

func (h *MedicineHandler) GetAllMedicines(c *gin.Context) {
	var filter models.MedicineFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	medicines, err := h.service.GetAllMedicines(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, medicines)
}

func (h *MedicineHandler) GetMedicineByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid medicine id",
		})
		return
	}

	medicine, err := h.service.GetMedicineByID(uint(id))
	if err != nil {
		if errors.Is(err, service.ErrMedicineNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHandler) UpdateMedicine(c *gin.Context) {
	var req models.MedicineUpdateRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid medicine id",
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	medicine, err := h.service.UpdateMedicine(uint(id), req)
	if err != nil {
		if errors.Is(err, service.ErrMedicineNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, service.ErrPriceMustBePositive) ||
			errors.Is(err, service.ErrStockQuantityNegative) {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHandler) DeleteMedicine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid medicine id",
		})
		return
	}

	if err := h.service.DeleteMedicine(uint(id)); err != nil {
		if errors.Is(err, service.ErrMedicineNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}