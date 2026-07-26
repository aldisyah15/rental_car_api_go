package handler

import (
	"log"
	"net/http"
	_ "rental_car/internal/car/model"
	model "rental_car/internal/car/model"
	"rental_car/internal/car/useCase"
	"rental_car/internal/platform"

	"github.com/gin-gonic/gin"
)

type HandlerCar struct {
	UseCaseCar *useCase.CarUseCase
}

func NewHandlerCar(car *useCase.CarUseCase) *HandlerCar {
	return &HandlerCar{
		UseCaseCar: car,
	}
}

func (hc HandlerCar) UploadRentalCar(c *gin.Context) {
	getRole, _ := c.Get("rule")
	rule := getRole.(string)

	if rule != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you don't have access",
		})
		return
	}

	var req = model.RequestCar{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorr": err.Error()})
		return
	}

	log.Printf("req : %v", req)

	form, err := c.MultipartForm()

	file := form.File["images"]
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_fileHeader": err.Error(),
		})
		return
	}

	var images []string

	for _, fileHeader := range file {
		url, err := platform.UploadFile(c.Request.Context(), fileHeader, "rental-car-app")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_inUploadFile": err.Error(),
			})
			return
		}
		images = append(images, url)
	}

	log.Printf("images : %v", images)
	//imageUrl, errUploadFile := platform.UploadFile(c.Request.Context(), fileHeader, "rental-car-app")
	//if errUploadFile != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error_inUploadFile": err.Error(),
	//	})
	//	return
	//}

	err = hc.UseCaseCar.UploadRentalCar(&req, images)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_UploadRentalCar": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
