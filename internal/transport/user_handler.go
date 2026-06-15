package transport

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", h.Register)
		users.GET("/:id", h.GetByID)
		
	}
}


func (h *UserHandler) Register(c *gin.Context) {
	var req *models.UserCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Неверный формат запроса",	
		})
	}

	err := h.userService.Create(req)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Пользователь с таким email уже существует",
			})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно зарегистрирован",
	})

}



func (h *UserHandler) GetByID(c *gin.Context) {
	// Получаем параметр id из строки запроса
	idStr := c.Param("id")

	
	
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат ID пользователя. Ожидается положительное число",
		})
		return
	}

	// 3. Вызов бизнес-логики через сервис
	user, err := h.userService.GetById(uint(id))
	if err != nil {
		
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Пользователь с таким ID не найден",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Внутренняя ошибка сервера при получении данных",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

