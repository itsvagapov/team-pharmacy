package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)

type ReviewHandler struct {
	service service.ReviewService
}

func NewReviewHandler(service service.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) RegisterRoutes(r *gin.Engine) {
	reviews := r.Group("/reviews")
	{
		reviews.PATCH("/:id", h.UpdateReview)
		reviews.DELETE("/:id", h.DeleteReview)
	}

	medicines := r.Group("/medicines/:id/reviews")
	{
		medicines.GET("", h.GetReviewsByMedicineID)
		medicines.POST("", h.CreateReview)
	}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req models.ReviewCreateRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid medicine id",
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.CreateReview(uint(id), req)
	if err != nil {
		if errors.Is(err, service.ErrRatingOutOfRange) ||
			errors.Is(err, service.ErrReviewTextRequired) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

func (h *ReviewHandler) GetReviewsByMedicineID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid medicine id",
		})
		return
	}

	reviews, err := h.service.GetReviewsByMedicineID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	var req models.ReviewUpdateRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid review id",
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	review, err := h.service.UpdateReview(uint(id), req)
	if err != nil {
		if errors.Is(err, service.ErrReviewNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		if errors.Is(err, service.ErrReviewTextRequired) {
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

	c.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid review id",
		})
		return
	}

	if err := h.service.DeleteReview(uint(id)); err != nil {
		if errors.Is(err, service.ErrReviewNotFound) {
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
