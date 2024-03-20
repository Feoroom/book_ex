package main

import (
	"book_ex/cmd/web"
	"book_ex/internal/models"
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {

	//Command-line flags
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	//Logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//DB connection
	connStr := "user=postgres dbname=webapp password=123 sslmode=disable"
	//connStr := "postgres://postgres:123@localhost/webapp?sslmode=disable"
	db, err := openDB(connStr)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	//Application
	app := &web.Application{
		ErrorLog: infoLog,
		InfoLog:  errorLog,
		Items:    &models.ItemModel{DB: db},
	}

	//Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	infoLog.Printf("Starting server on: %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
