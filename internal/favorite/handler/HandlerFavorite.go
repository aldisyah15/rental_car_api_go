package handler

import (
	"net/http"
	"rental_car/internal/favorite/model"
	"rental_car/internal/favorite/useCase"

	"github.com/gin-gonic/gin"
)

type HandlerFavorite struct {
	useCase *useCase.FavoriteUseCase
}

func NewHandlerFavorite(useCase *useCase.FavoriteUseCase) *HandlerFavorite {
	return &HandlerFavorite{useCase: useCase}
}

func (h HandlerFavorite) Favorite(c *gin.Context) {
	value, _ := c.Get("userName")
	username := value.(string)
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "StatusUnauthorized",
		})
		return
	}
	var req model.RequestFavorite

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	favorite := &model.RequestFavorite{
		Username: username,
		IdCar:    req.IdCar,
	}

	err = h.useCase.AddAndRemoveFavorite(favorite)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h HandlerFavorite) GetFavorite(c *gin.Context) {
	value, _ := c.Get("userName")
	if value == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorize",
		})
		return
	}
	username := value.(string)

	result, err := h.useCase.GetFavorite(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}
