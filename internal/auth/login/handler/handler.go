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
		c.JSON(http.StatusBadGateway, gin.H{
			"Error": "gagal masuk",
		})
	}
	user, err2 := h.useCase.AuthLogin(&req)
	if err2 != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"Error": err2.Error(),
		})
	}

	tokenString, _ := middleware.GenerateString(user.UserName)
	c.JSON(http.StatusOK, gin.H{
		"Message": "success",
		"token":   tokenString,
	})

}
