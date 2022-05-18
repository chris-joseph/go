package api

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"todo/pkg/config"
)

type Middleware struct {
	config *config.Settings
}

func (m Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Missing Authorization Header"})
			c.Abort()
		}
		type Claims struct {
			Id  string `json:"id"`
			Exp int    `json:"exp"`
			jwt.StandardClaims
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.config.JwtSecret), nil
		})
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "UnAuthorized"})
		}
		claims, ok := token.Claims.(*Claims)
		if ok && token.Valid {
			c.Set("user", claims.Id)
			c.Next()
		} else {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "UnAuthorized"})
		}
	}

}
