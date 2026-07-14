package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func GenerateString(userName string) (string, error) {
	claim := jwt.MapClaims{
		"userId": userName,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(secretKey)
}

func veryfyToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	userId, ok := claims["userId"].(string)
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err == nil && token.Valid {
		if ok == true {
			return userId, nil
		}
	}

	return "", err
}

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		userId, err := veryfyToken(strings.TrimPrefix(authorization, "Bearer "))

		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		ctx := context.WithValue(req.Context(), "userId", userId)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
