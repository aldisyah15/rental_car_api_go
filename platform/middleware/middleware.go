package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func GenerateString(userName string, rule string) (string, error) {
	claim := jwt.MapClaims{
		"username": userName,
		"rule":     rule,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(secretKey)
}

func veryfyToken(tokenString string) (string, string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	username, ok := claims["username"].(string)
	rule, ok := claims["rule"].(string)

	if err == nil && token.Valid {
		if ok == true {
			return username, rule, nil
		}
	}

	return "", "", err
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		userName, rule, err := veryfyToken(strings.TrimPrefix(authorization, "Bearer "))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak memiliki akses"})
			return
		}

		log.Printf("log dari middleware: %v", userName)

		c.Set("userName", userName)
		c.Set("rule", rule)
		c.Next()
	}
}
