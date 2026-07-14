package handler

import (
	"net/http"
	"rental_car/internal/auth/login/model"
	useCase2 "rental_car/internal/auth/login/useCase"
	"rental_car/internal/platform/middleware"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	useCase *useCase2.Login
}

func NewLoginHandler(uc *useCase2.Login) *LoginHandler {
	return &LoginHandler{
		useCase: uc,
	}
}

func (h *LoginHandler) LoginUser(c *gin.Context) {
	var req model.RequestLogin

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Token Invalid",
		})
		return
	}

	user, _ := h.useCase.AuthLogin(&req)
	tokenString, err := middleware.GenerateString(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Token Invalid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"data":    user,
		"tokem":   tokenString,
	})
}
