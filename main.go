package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/beinganukul/rest_api_golang/controllers/auth"
	"github.com/beinganukul/rest_api_golang/controllers/book"
	"github.com/beinganukul/rest_api_golang/controllers/homepage"
	"github.com/beinganukul/rest_api_golang/initializers"
	"github.com/beinganukul/rest_api_golang/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnVars()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/", homepage.Home)
	r.GET("/search", book.SearchBooks)

	r.POST("/signup", auth.Signup)
	r.POST("/login", auth.Login)
	r.POST("/change_password", middleware.RequireAuth, auth.ChangePassword)
	r.POST("/upload", middleware.RequireAuth, book.UploadBook)
	r.POST("/homepage/reindex", homepage.ReindexImage)

	mediaPath, _ := filepath.Abs(os.Getenv("MEDIA_PATH"))
	r.StaticFS("/media", http.Dir(mediaPath))
	r.Run()
}
