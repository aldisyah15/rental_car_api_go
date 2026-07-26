package handler

import (
	"log"
	"net/http"
	"rental_car/internal/auth/register/model"
	"rental_car/internal/auth/register/useCase"

	"github.com/gin-gonic/gin"
)

type Register struct {
	AuthCase *useCase.AuthCase
}

func NewRegisterHandler(uc *useCase.AuthCase) *Register {

	return &Register{
		AuthCase: uc,
	}
}

func (h *Register) CreateUser(c *gin.Context) {
	var req model.RegisterRequest
	err := c.ShouldBindJSON(&req)
	log.Print(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.AuthCase.CreateUser(&req)
	log.Printf("handler err: %v", err)

	if err != nil {
		c.JSON(http.StatusFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "success",
		})
	}

}
