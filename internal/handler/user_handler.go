package handler

import (
	"food-delivery-api/constants/user_role"
	"food-delivery-api/internal/service"
	"food-delivery-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Name     string             `json:"name"`
		Email    string             `json:"email"`
		Password string             `json:"password"`
		UserRole user_role.UserRole `json:"userRole"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Register(req.Name, req.Email, req.Password, req.UserRole); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusCreated, "user registered")
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		utils.JSONError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, gin.H{"accessToken": accessToken})
}
