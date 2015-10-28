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
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/audio", "audio")
	router.Static("/css", "css")
	router.Static("/images", "images")
	router.Static("/js", "js")
	router.Static("/video", "video")


	router.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK, "We're getting married. Give us money.")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/loaderio-21e121865e59f3867a444fbdb50f665d/", func(c *gin.Context) {
		c.String(http.StatusOK, "loaderio-21e121865e59f3867a444fbdb50f665d")
	})

	router.Run(":" + port)
}
