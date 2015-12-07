package main

import (
	"database/sql"
	"jakerandell.com/wedding/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"jakerandell.com/wedding/Godeps/_workspace/src/github.com/gorilla/sessions"
	_ "jakerandell.com/wedding/Godeps/_workspace/src/github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var (
	db            *sql.DB = nil
	validPassword         = os.Getenv("THE-PASSWORD")
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

	store := sessions.NewCookieStore([]byte(os.Getenv("COOKIE-STORE-SECRET")))

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/css", "css")
	router.Static("/images", "images")
	router.Static("/js", "js")
	router.Static("/fonts", "fonts")

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.POST("/login", func(c *gin.Context) {
		pass := c.PostForm("password")

		session, err := store.Get(c.Request, "gatekeeper")
		if err != nil {
			c.Error(err)
			return
		}

		if pass == validPassword {
			session.Values["isLoggedIn"] = "sure"
			session.Save(c.Request, c.Writer)
			c.Redirect(http.StatusFound, "/address")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"success": "false",
				"message": "Incorrect Password",
			})
		}
	})
	router.Use(validateAuth(store))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
//		c.Redirect(http.StatusTemporaryRedirect, "/address")
	})

	router.GET("/address", func(c *gin.Context) {
		c.HTML(http.StatusOK, "address-form.html", gin.H{})
	})

	router.POST("/address-post", func(c *gin.Context) {
		name := c.PostForm("name")
		addr1 := c.PostForm("addr1")
		addr2 := c.PostForm("addr2")
		city := c.PostForm("city")
		state := c.PostForm("state")
		zip := c.PostForm("zip")
		phone := c.PostForm("phone")

		if _, err := db.Query(`INSERT INTO addresses
			(name, addr1, addr2, city, state, zip, phone, origination)
			VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`,
			name, addr1, addr2, city, state, zip, phone,
		); err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"error":   err,
			})
		} else {
			c.JSON(200, gin.H{
				"success": true,
			})
		}
	})

	router.GET("/loaderio-21e121865e59f3867a444fbdb50f665d/", func(c *gin.Context) {
		c.String(http.StatusOK, "loaderio-21e121865e59f3867a444fbdb50f665d")
	})

	router.Run(":" + port)

}

func validateAuth(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "gatekeeper")
		if err != nil {
		}
		if session.Values["isLoggedIn"] == "sure" {
			c.Next()
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
	}
}
