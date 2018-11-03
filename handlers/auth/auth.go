package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

const secret = "hahaha"

func CheckAuth(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return secret, nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	hecan, ok := claims["valid"].(bool)
	if !ok || !hecan {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}

	c.Next()
}
