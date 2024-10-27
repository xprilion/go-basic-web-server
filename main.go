package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/greet", func(c *gin.Context) {
		name := c.PostForm("name")
		c.HTML(http.StatusOK, "greet.html", gin.H{"name": name})
	})

	router.Run(":8080")
}
