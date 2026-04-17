package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func home(c *gin.Context){
	c.String(http.StatusOK, "We are home")
}

func viewSnippet(c *gin.Context) {
	idx := Query("id")

	err, id := strconv.Atoi(idx)
	if err != nil || id < 1 {
		c.String(http.StatusNotFound, "Snippet not found")
		return
	}
}