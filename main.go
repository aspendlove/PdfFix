package main

import (
	"log"
	"pdf-fix/src/config"
	"pdf-fix/src/pdf"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Printf("Failed to load config file, falling back to defaults\n")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.POST("/api/fix-pdf", pdf.SubmitHandler(cfg))
	router.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})
	router.Run()
}
