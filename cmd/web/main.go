package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	
	// Define an application struct to hold the application-wide dependencies for the
	// web application. For now we'll only include fields for the two custom loggers, but
	// we'll add more to it as the build progresses.

	type application struct{
		errorLog *log.Logger,
		infoLog *log.Logger
	}

func main () {

	// a new command line flag is defined
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// Creating a new error log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for written error message
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Instead of using Default we use NEW
	r := gin.New()
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("./ui/html/pages/*")
	
	// This part serve the static files also remove the prefix
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

	app := &application{
		errorLog: errorLog
		infoLog: infoLog
	}

	// add routers
	r.GET("/", app.home)
	r.GET("/snippet/view", app.viewSnippet)
	r.Any("/snippet/create", app.createSnippet)

	// Initializing a new server Struct
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: r
	}

	infoLog.Printf("Server is running on Port %s:", *addr)
	err := srv.Run()
	errorLog.Fatal(err)
}
