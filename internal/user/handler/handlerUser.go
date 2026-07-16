package handler

import (
	"net/http"
	"rental_car/internal/user/useCase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useCase *useCase.UserUseCase
}

func NewUserHandler(u *useCase.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: u,
	}
}

func (u UserHandler) GetUser(c *gin.Context) {
	getUsername, _ := c.Get("userName")

	username := getUsername.(string)
	user, err := u.useCase.GetUser(username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if user != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}

}
