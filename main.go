package main

import (
//	"jakerandell.com/wedding/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"net/http"
	"os"
	"github.com/gorilla/mux"
"html/template"
	"log"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		//		log.Fatal("$PORT must be set")
		port = "5000"
	}

//	router := gin.New()
//	router.Use(gin.Logger())
//	router.LoadHTMLGlob("templates/*.html")
//	router.Static("/audio", "audio")
//	router.Static("/css", "css")
//	router.Static("/images", "images")
//	router.Static("/js", "js")
//	router.Static("/video", "video")


	/*router.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK, "We're getting married. Give us money.")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})*/

	/*router.GET("/loaderio-21e121865e59f3867a444fbdb50f665d/", func(c *gin.Context) {
		c.String(http.StatusOK, "loaderio-21e121865e59f3867a444fbdb50f665d")
	})*/

//	router.Run(":" + port)


	r := mux.NewRouter()
//	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	staticHandler := http.FileServer(http.Dir("."))
	staticHandler = http.StripPrefix("/static/", staticHandler)
	r.PathPrefix("/static/").Handler(staticHandler)

	tmpl := template.Must(template.ParseGlob("templates/*"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

//	r.Handle("/", r)
	log.Fatal(http.ListenAndServe(":" + port, r))
}
