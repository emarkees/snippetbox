package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func home(c *gin.Context){
	c.String(http.StatusOK, "We are home")
}