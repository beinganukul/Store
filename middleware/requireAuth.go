package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/beinganukul/rest_api_golang/initializers"
	"github.com/beinganukul/rest_api_golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// get cookie from request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// decode/validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// checks if the token has not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token has expired!",
			})
		}
		var user models.User
		initializers.DB.First(&user, claims["user"])

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token!",
			})
		}
		// attaches the user object to the request
		c.Set("user", user)

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
