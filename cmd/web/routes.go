package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

func (app *application) routes() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.StaticFS("/static", http.Dir("./ui/static"))

	// Custom 404 for unmatched routes
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
			"message": "Page not found",
		})
		// or: c.String(http.StatusNotFound, "Not Found")
	})

	r.NoMethod(func(c *gin.Context) {
		c.Header("Allow", "GET, POST")
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	})

	// add routers
	r.GET("/", app.home)
	r.GET("/snippet/view", app.viewSnippet)
	r.POST("/snippet/create", app.createSnippet)

	return r
}