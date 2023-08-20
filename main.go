package main

import (
	"github.com/beinganukul/rest_api_golang/controllers/auth"
	"github.com/beinganukul/rest_api_golang/controllers/book"
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

	//r.POST("/posts", controllers.CreatePosts)
	//r.GET("/posts", controllers.ListPosts)
	//r.GET("/post/:id", controllers.GetPost)
	//r.PUT("/post/:id", controllers.UpdatePost)
	//r.DELETE("/post/:id", controllers.DeletePost)
	r.POST("/signup", auth.Signup)
	r.POST("/login", auth.Login)
	r.POST("/change_password", middleware.RequireAuth, auth.ChangePassword)
	r.POST("/upload", middleware.RequireAuth, book.UploadBook)
	r.Run()
}
