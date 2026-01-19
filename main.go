package main

import (
	"pdf-fix/src/pdf"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.POST("/api/fix-pdf", pdf.SubmitHandler)
	router.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})
	router.Run()
}
