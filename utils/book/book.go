package book_util

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func SaveBookImageToDisk(c *gin.Context) ([]string, error) {
	var images []string
	form, err := c.MultipartForm()

	if err != nil {
		return nil, err
	}
	files := form.File["Image"]

	absolutePath, _ := filepath.Abs((os.Getenv("MEDIA_PATH") + "book"))

	for _, file := range files {

		filename := ulid.Make().String() + "_" + file.Filename
		images = append(images, filename)

		dst := filepath.Join(absolutePath, filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			return nil, err
		}
	}
	return images, err
}

func CalculateSlug(title string) string {
	regex := regexp.MustCompile("\\W")
	Slug := regex.ReplaceAllString(strings.ToLower(title), "-")
	Slug = strings.Trim(Slug, "-")
	return Slug
}
