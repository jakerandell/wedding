package main

import (
	"database/sql"
	"jakerandell.com/wedding/Godeps/_workspace/src/github.com/gin-gonic/gin"
	_ "jakerandell.com/wedding/Godeps/_workspace/src/github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var (
	db *sql.DB = nil
)

func main() {
	port := os.Getenv("PORT")
	var errd error

	if port == "" {
		//		log.Fatal("$PORT must be set")
		port = "5000"
	}

	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/css", "css")
	router.Static("/images", "images")
	router.Static("/js", "js")
	router.Static("/fonts", "fonts")

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"randellwedding": "august13",
	}))

	router.GET("/", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.html", gin.H{})
		c.Redirect(http.StatusTemporaryRedirect, "/address")
	})

	authorized.GET("/address", func(c *gin.Context) {
		c.HTML(http.StatusOK, "address-form.html", gin.H{})
	})

	router.POST("/address-post", func(c *gin.Context) {
		name := c.PostForm("q1")
		addr1 := c.PostForm("q2")
		addr2 := c.PostForm("q3")
		city := c.PostForm("q4")
		state := c.PostForm("q5")
		zip := c.PostForm("q6")
		phone := c.PostForm("q7")

		if _, err := db.Query(`INSERT INTO addresses
			(name, addr1, addr2, city, state, zip, phone, origination)
			VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`,
			name, addr1, addr2, city, state, zip, phone,
		); err != nil {
			//			c.String(http.StatusInternalServerError, fmt.Sprintf("Error incrementing tick: %q", err))
			c.JSON(200, gin.H{
				"success": false,
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"success": true,
			})
		}

		/*c.JSON(200, gin.H{
			"message": name,
			"address 1": addr1,
			"address 2": addr2,
			"city": city,
			"state": state,
			"zip": zip,
			"phone": phone,
		})*/
	})

	router.GET("/loaderio-21e121865e59f3867a444fbdb50f665d/", func(c *gin.Context) {
		c.String(http.StatusOK, "loaderio-21e121865e59f3867a444fbdb50f665d")
	})

	router.Run(":" + port)

}
