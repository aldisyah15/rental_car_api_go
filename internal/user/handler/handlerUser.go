package handler

import (
	"fmt"
	"log"
	"net/http"
	"rental_car/internal/user/model"
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
	getRule, _ := c.Get("rule")

	username := getUsername.(string)
	rule := getRule.(string)
	log.Printf("rule: %v", rule)
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

func (u UserHandler) UpdateUser(c *gin.Context) {
	getUsername, _ := c.Get("userName")
	if getUsername == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userNotfound",
		})
	}
	log.Printf("loghandlerUser: %v", getUsername)
	var req = model.RequestUpdateUser{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	username := getUsername.(string)
	log.Printf("log_username: %v", username)
	updateUseCase, err := u.useCase.UpdateUser(username, req)
	log.Printf("log_updateUseCase: %v", updateUseCase)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	if updateUseCase != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": updateUseCase,
		})
	}
}

func (u UserHandler) DeleteUser(c *gin.Context) {
	getUsername, _ := c.Get("userName")

	user := getUsername.(string)
	result, err := u.useCase.DeleteUser(user)
	fmt.Println("Result:", result, "Error:", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("id %v has been delete from database", result),
	})
}
