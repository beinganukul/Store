package homepage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	var homeData struct {
		Banners []string `json:"banners"`
	}
	file, fileStatus := filepath.Abs(os.Getenv("CRM_DATA_FILE"))

	if fileStatus != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to read homepage data!",
		})
		return
	}
	crm, _ := os.ReadFile(file)
	_ = json.Unmarshal([]byte(crm), &homeData)
	fmt.Println(homeData.Banners)
}

func ReindexImage(c *gin.Context) {
	var reindexedImage struct {
		Banners []string
	}
	c.Bind(&reindexedImage)

	var openedData struct {
		Banners []string `json:"banners"`
	}
	file, fileStatus := filepath.Abs(os.Getenv("CRM_DATA_FILE"))

	if fileStatus != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to read homepage data!",
		})
		return
	}
	opnData, _ := os.ReadFile(file)
	_ = json.Unmarshal([]byte(opnData), &openedData)

	openedData.Banners = reindexedImage.Banners

	updatedData, _ := json.Marshal(openedData)

	err := os.WriteFile(file, updatedData, 0644)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to write file!",
		})
		return
	}

}
