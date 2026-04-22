package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"html/template"
)

func (app *application) home(c *gin.Context) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = ts.ExecuteTemplate(c.Writer, "base", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
}

func (app *application)viewSnippet(c *gin.Context) {
	idx := c.Query("id")
	id, err := strconv.Atoi(idx)
	if err != nil || id < 1 {
		c.String(http.StatusNotFound, "Snippet not found")
		return
	}
	c.String(http.StatusOK, "Snippet ID: %d", id)
}

func (app *application) createSnippet(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.Header("Allow", http.MethodPost)
		c.String(http.StatusMethodNotAllowed, "Method not Allowed")
		return
	}

	c.String(http.StatusOK, "Snippet created successfully")
}