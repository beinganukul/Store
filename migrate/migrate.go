package main

import (
	"github.com/beinganukul/rest_api_golang/initializers"
	"github.com/beinganukul/rest_api_golang/models"
)

func init() {
	initializers.LoadEnVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
	initializers.DB.AutoMigrate(&models.User{})
}
