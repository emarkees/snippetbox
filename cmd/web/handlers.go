package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	// "html/template"
	"github.com/emarkees/snippetbox/internal/models"

)

func (app *application) home(c *gin.Context) {
	snippets, err := app.snippet.Latest()
	if err != nil {
		app.serverError(c, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(c.Writer, "%+v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(c, err)
	// 	return
	// }

	// err = ts.ExecuteTemplate(c.Writer, "base", nil)
	// if err != nil {
	// 	app.serverError(c, err)
	// 	return
	// }
}

func (app *application) createSnippet(c *gin.Context) {
	
	title := "O Grace"
	content := "O Grace of the Most high\nClimb mount Fuji,\nBut slow, slowly!\n\n- Kabiyyoshi Issa"
	expires := 9

	id, err := app.snippet.Insert(title, content, expires)
	if err != nil {
		app.serverError(c, err)
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/snippet/view?id=%d\n", id))
}

func (app *application)viewSnippet(c *gin.Context) {
	idx := c.Query("id")
	id, err := strconv.Atoi(idx)
	if err != nil || id < 1 {
		app.notFound(c)
		return
	}

	snippet, err := app.snippet.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(c)
		} else {
			app.serverError(c, err)
		}
		return
	}
	fmt.Fprintf(c.Writer, "%+v", snippet)
}
