package controllers

//func CreatePosts(c *gin.Context) {
//	// get data from request body and create a post
//
//	var body struct {
//		Body  string
//		Title string
//	}
//
//	c.Bind(&body)
//
//	post := models.Post{Title: body.Title, Body: body.Body}
//	result := initializers.DB.Create(&post)
//
//	if result.Error != nil {
//		c.Status(400)
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"post": post,
//	})
//}
//
//func ListPosts(c *gin.Context) {
//	// list out all the data from posts table
//
//	var posts []models.Post
//	initializers.DB.Find(&posts)
//
//	c.JSON(200, gin.H{
//		"posts": posts,
//	})
//}
//
//func GetPost(c *gin.Context) {
//	// get specific post using it's id
//	id := c.Param("id")
//
//	var post models.Post
//	initializers.DB.First(&post, id)
//
//	c.JSON(200, gin.H{
//		"post": post,
//	})
//}
//
//func UpdatePost(c *gin.Context) {
//	// get the request body, get the specific post using id from url, and update the post
//
//	var body struct {
//		Title string
//		Body  string
//	}
//
//	c.Bind(&body)
//
//	id := c.Param("id")
//
//	var post models.Post
//	initializers.DB.First(&post, id)
//
//	initializers.DB.Model(&post).Updates(
//		models.Post{
//			Title: body.Title,
//			Body:  body.Body,
//		},
//	)
//	c.JSON(200, gin.H{
//		"post": post,
//	})
//}
//
//func DeletePost(c *gin.Context) {
//	// select post and delete
//	id := c.Param("id")
//
//	initializers.DB.Delete(&models.Post{}, id)
//
//	c.Status(200)
//}
//
