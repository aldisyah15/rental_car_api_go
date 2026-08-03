package handler

import (
	"net/http"
	"rental_car/internal/booking/model"
	"rental_car/internal/booking/useCase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	useCase *useCase.BookingUseCase
}

func NewHandlerBooking(useCase *useCase.BookingUseCase) *BookingHandler {
	return &BookingHandler{
		useCase: useCase,
	}
}

func (h BookingHandler) ProsesToCheckout(c *gin.Context) {
	getUsername, _ := c.Get("userName")
	if getUsername == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorize",
		})
		return
	}

	username := getUsername.(string)
	result := model.RequestProsesToCheckout{
		Username: username,
	}
	err := c.ShouldBindJSON(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := h.useCase.ProsesToCheckout(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (h BookingHandler) GetDetailCheckout(c *gin.Context) {
	getUsername, _ := c.Get("userName")
	if getUsername == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorize",
		})
		return
	}
	username := getUsername.(string)
	param := c.Params.ByName("idOrder")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := h.useCase.GetDetailCheckout(id, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    result,
	})

}

func (h BookingHandler) BookingNow(c *gin.Context) {
	getUsername, _ := c.Get("userName")
	if getUsername == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorize",
		})
		return
	}
	idOrder := c.Params.ByName("idOrder")
	res, err := h.useCase.BookingNow(idOrder, getUsername.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
	return
}
