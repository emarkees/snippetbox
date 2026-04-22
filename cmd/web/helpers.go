package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func (app *application) serverError(c *gin.Context, err error) {
	if err == nil {
		app.errorLog.Println("serverError called with nil error")
		c.Status(http.StatusInternalServerError)
		return
	}

	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace) // skip 2 frames to show caller, not this helper

	c.Status(http.StatusInternalServerError)
}

func (app *application) clientError(c *gin.Context, status int) {
	if status < 100 || status > 599 {
		app.errorLog.Printf("invalid HTTP status: %d", status)
		status = http.StatusInternalServerError
	}
	c.Status(status)
}

func (app *application) notFound(c *gin.Context) {
	app.clientError(c, http.StatusNotFound)
}