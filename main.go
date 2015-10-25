package main

import (
	"jakerandell.com/wedding/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		//		log.Fatal("$PORT must be set")
		port = "5000"
	}

	router := gin.New()
	router.Use(gin.Logger())
//	router.LoadHTMLGlob("templates/*.tmpl.html")
//	router.Static("/static", "static")


	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "We're getting married. Give us money.")
	})

	router.Run(":" + port)
}
