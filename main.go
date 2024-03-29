package main

import (
	"book_ex/cmd/web"
	"book_ex/internal/models"
	"crypto/tls"
	"database/sql"
	"flag"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

// TODO: createReview Session Store
// go run C:\Users\ezioe\sdk\go1.21.4\src\crypto\tls\generate_cert.go --rsa-bits=2048 --host=localhost
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

	//Initializing the template cache
	templCache, err := web.NewTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	//Form decoder

	formDecoder := form.NewDecoder()

	//Session manager
	sessionManager := scs.New()
	//sessionManager.Store = pqstore.New(db)
	sessionManager.Lifetime = time.Hour
	sessionManager.Cookie.Secure = true

	//
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	//Application
	app := &web.Application{
		ErrorLog:       errorLog,
		InfoLog:        infoLog,
		Reviews:        &models.ReviewModel{DB: db},
		Users:          &models.UserModel{DB: db},
		Books:          &models.BookModel{DB: db},
		TemplateCache:  templCache,
		FormDecoder:    formDecoder,
		SessionManager: sessionManager,
	}

	//Server struct
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.Routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on: %s", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
