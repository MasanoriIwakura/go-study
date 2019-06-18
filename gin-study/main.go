package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Link struct {
	Name string
	Url  string
}

var links = []Link{
	Link{Name: "About", Url: "/about"},
}

const (
	base = "templates/layouts/base.html"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/css", "templates/css")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		html := template.Must(template.ParseFiles(base, "templates/index.html"))
		r.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.html", gin.H{
			"message": "Template Sample",
			"links":   links,
		})
	})

	r.GET("/about", func(c *gin.Context) {
		html := template.Must(template.ParseFiles(base, "templates/pages/about.html"))
		r.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.html", gin.H{})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
