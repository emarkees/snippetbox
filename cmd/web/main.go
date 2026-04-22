package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	// Creating a new error log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for written error message
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()


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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
