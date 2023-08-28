package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/beinganukul/rest_api_golang/initializers"
	"github.com/beinganukul/rest_api_golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	var body struct {
		Email    string
		Password string
		FName    string
		LName    string
		Mobile   int32
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password!",
		})
		return
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
		FName:    body.FName,
		LName:    body.LName,
		Mobile:   body.Mobile,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user!",
		})
		return
	}

	c.Status(201)
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password!",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	fmt.Println(err)
	fmt.Println(user.Password, user.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a token!",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"jwt": tokenString,
	})
}

func ChangePassword(c *gin.Context) {

	var body struct {
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	c.Bind(&body)

	if body.Password != body.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password doesn't match!",
		})
		return
	}

	authuser, userStatus := c.Get("user")
	if userStatus != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user doesn't exist!",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "couldn't hash the password!",
		})
		return
	}

	user := authuser.(models.User)
	user.Password = string(hash)
	result := initializers.DB.Model(&user).Update("Password", string(hash))
	fmt.Println(string(hash))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to update password!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password changed successfully!",
	})
}
