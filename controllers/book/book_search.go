package book

//func SearchBooks(c *gin.Context) {
//	type searchOutput struct {
//		Id     int32          `gorm:"column:id"`
//		Title  string         `gorm:"column:title"`
//		BImage pq.StringArray `gorm:"column:b_image"`
//		Rank   float32        `gorm:"column:rank"`
//		UserID int32          `gorm:"column:user_id"`
//		FName  string         `gorm:"column:f_name"`
//		LName  string         `gorm:"column:l_name"`
//		UImage pq.StringArray `gorm:"column:u_image"`
//	}
//
//	var books []searchOutput
//
//	searchString := c.Query("keyword")
//	fmt.Println(searchString)
//
//	// query to extract books
//	var query string = `
//	SELECT
//		b.image AS b_image,
//		b.rank,
//		b.user_id,
//		users.f_name,
//		users.l_name,
//		CASE
//			WHEN users.image IS NULL THEN ARRAY[]::TEXT[]
//			ELSE ARRAY[users.image]
//		END AS u_image
//	FROM
//		search_books(?) AS b
//	INNER JOIN
//		users ON b.user_id = users.id;
//	`
//
//	if err := initializers.DB.Raw(query, searchString).Scan(&books).Error; err != nil {
//		log.Fatal("Failed to execute the query!")
//	}
//	fmt.Println(books)
//}

import (
	"net/http"

	"github.com/beinganukul/rest_api_golang/initializers"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func SearchBooks(c *gin.Context) {
	type searchOutput struct {
		Id     int32          `gorm:"column:id"`
		Title  string         `gorm:"column:title"`
		BImage pq.StringArray `gorm:"column:b_image"`
		Rank   float32        `gorm:"column:rank"`
		UserID int32          `gorm:"column:user_id"`
		FName  string         `gorm:"column:f_name"`
		LName  string         `gorm:"column:l_name"`
		UImage string         `gorm:"column:u_image"`
	}
	var books []searchOutput

	searchString := c.Query("keyword")

	// query to extract books
	var query string = `
	SELECT 
		b.id AS id,
		b.title AS title,
		b.rank AS rank, 
		b.user_id AS user_id, 
		users.f_name AS f_name, 
		users.l_name AS l_name, 
		COALESCE(b.image, ARRAY[]::TEXT[]) AS b_image,
		users.image AS u_image 
	FROM 
		search_books(?) AS b
	INNER JOIN 
		users ON b.user_id = users.id;	
	`

	if err := initializers.DB.Raw(query, searchString).Scan(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't run query and bind data!",
		})
		return
	}

	c.JSON(http.StatusOK, books)
}
