package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		html := template.Must(template.ParseFiles("templates/layouts/base.html", "templates/index.html"))
		r.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.html", gin.H{
			"message": "Template Sample",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
