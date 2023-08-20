package book

import (
	"net/http"

	"github.com/beinganukul/rest_api_golang/initializers"
	"github.com/beinganukul/rest_api_golang/models"
	book_util "github.com/beinganukul/rest_api_golang/utils/book"
	"github.com/gin-gonic/gin"
)

func UploadBook(c *gin.Context) {
	authUser, authUserStatus := c.Get("user")

	if authUserStatus != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found!",
		})
		return
	}
	user := authUser.(models.User)

	// user generated
	var formBody struct {
		Title    string
		Author   []string
		ISBN     string
		Weight   int
		Language string
		Pages    int
		Price    float32
		Genre    []string
	}
	err := c.Bind(&formBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid form data!",
		})
		return
	}

	var manuallyDefined struct {
		Slug      string
		Available bool
		Images    []string
		UserID    uint
	}

	// NON User Generated
	// generating slug
	manuallyDefined.Slug = book_util.CalculateSlug(formBody.Title)
	manuallyDefined.Available = true
	manuallyDefined.UserID = user.ID
	manuallyDefined.Images, err = book_util.SaveBookImageToDisk(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to save file!",
		})
		return
	}

	finalData := models.Book{
		Title:     formBody.Title,
		Author:    formBody.Author,
		ISBN:      formBody.ISBN,
		Weight:    formBody.Weight,
		Language:  formBody.Language,
		Pages:     formBody.Pages,
		Price:     formBody.Price,
		Genre:     formBody.Genre,
		Slug:      manuallyDefined.Slug,
		Available: manuallyDefined.Available,
		Image:     manuallyDefined.Images,
		UserID:    manuallyDefined.UserID,
	}

	result := initializers.DB.Create(&finalData)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "couldn't create a book instance!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"book": finalData,
	})
}
