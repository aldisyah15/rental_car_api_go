package handler

import (
	"log"
	"net/http"
	useCase2 "rental_car/internal/brand/useCase"
	"rental_car/platform"

	"github.com/gin-gonic/gin"
)

type LogoHandler struct {
	useCase *useCase2.LogoUseCase
}

func NewCarHandlerLogo(carUseCase *useCase2.LogoUseCase) *LogoHandler {
	return &LogoHandler{
		useCase: carUseCase,
	}
}

func (hc LogoHandler) UploadLogo(c *gin.Context) {
	rule, _ := c.Get("rule")
	if rule != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "StatusUnauthorized",
		})
		return
	}
	name := c.PostForm("name")
	err := c.ShouldBindJSON(&name)
	fileName, _ := c.FormFile("logo")

	url, err := platform.UploadFileAndGet(c.Request.Context(), fileName, "rental-car-app", "logo")
	log.Printf("url : %v", url)

	err = hc.useCase.UploadLogo(url, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (hc LogoHandler) GetAllLogo(c *gin.Context) {
	result, err := hc.useCase.GetAllLogo()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}
