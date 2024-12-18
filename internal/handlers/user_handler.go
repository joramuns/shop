package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/joramuns/shop/internal/models"
	us "github.com/joramuns/shop/internal/service"
	"net/http"
)

type UserHandler struct {
	service *us.UserService
}

func NewUserHandler(service *us.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdUser, err := h.service.RegisterUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
