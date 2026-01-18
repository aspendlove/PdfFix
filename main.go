package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.POST("/submit", func(c *gin.Context) {
		file, err := c.FormFile("pdf")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		fmt.Printf("Filename: %s, Size: %d bytes", file.Filename, file.Size)

		dst := "./uploads/" + file.Filename
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, "File %s uploaded successfully", file.Filename)
	})
	router.Run()
}
