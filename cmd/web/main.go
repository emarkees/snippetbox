package main

import (
	"flag"
	"log"
	"net/http"
	"os"
 	)
	
	// Define an application struct to hold the application-wide dependencies for the
	// web application. For now we'll only include fields for the two custom loggers, but
	// we'll add more to it as the build progresses.

	type application struct {
		errorLog *log.Logger
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

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	r := app.routes()

	// Initializing a new server Struct
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: r,
	}

	infoLog.Printf("Server is running on Port %s:", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
