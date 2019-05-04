package api

import (
	scanner "../scanner"
	"github.com/gin-gonic/gin"
	"net/http"

)

func getBookPathsInFolder(c *gin.Context) {
	folderPath := c.Query("path")
	paths, err := scanner.ScanFolder(folderPath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "",
			"data": []
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "",
		"data": paths,
	})
}

// RegisterBookInterfaces shall registe a series of interfaces for managing books and retrivering books' info
func RegisterBookInterfaces(r *gin.Engine) {
	bookRouter := r.Group("/book")
	{
		bookRouter.GET("/findbooks", getBookPathsInFolder)
	}
}