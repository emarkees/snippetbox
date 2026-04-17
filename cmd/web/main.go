package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main () {
	r := gin.Default()

	// no Route
	r.NoRoute(func(c *gin.Context){
		c.String(http.StatusNotFound, "404 page not found")
	})

	r.NoMethod(func(c *gin.Context) {
		c.String(http.StatusMethodNotFound, "405 Method not Allowed")
	})

	// add routers
	r.GET("/", home)

	log.Println("Server is running on Port: 4000")
	err := r.Run()
	log.Fatal(err)
}
