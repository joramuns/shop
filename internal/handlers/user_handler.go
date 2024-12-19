package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/joramuns/shop/internal/models"
	us "github.com/joramuns/shop/internal/service"
	"log"
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
		log.Println("Handler invalid input", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdUser, err := h.service.RegisterUser(c.Request.Context(), &user)
	if err != nil {
		log.Println("Handler error while registering user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers(c)
	if err != nil {
		log.Println("Handler error while listing users:", err)
	}

	c.JSON(http.StatusOK, users)
}
