package handler

import (
	"net/http"
	"rental_car/internal/auth/login/model"
	useCase2 "rental_car/internal/auth/login/useCase"
	"rental_car/platform/middleware"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	user, err := h.useCase.AuthLogin(&req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error(),
		})
		return
	}

	tokenString, _ := middleware.GenerateString(user.UserName, user.Rule)
	c.JSON(http.StatusOK, gin.H{
		"Message": "success",
		"token":   tokenString,
	})

}
