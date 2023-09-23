package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/lib/pq"

	"github.com/dhinogz/finch/internal/data"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
	models         data.Models
	config         config
}

// Initialize a new HTTP server for a web application
func main() {
	// Get environment variables
	var cfg config
	cfg, err := loadEnv()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	// Configure logger
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	infoLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)

	// Open database connection with config variables
	db, err := openDB(cfg)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	gmap, err := openGMap(cfg)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize model struct for db transaction functions
	models := data.NewModel(db, gmap)

	// Initialize templates into memory
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	//
	formDecoder := form.NewDecoder()

	//
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	// Initialize application struct for easy access to data and functions
	app := &application{
		config:         cfg,
		errorLog:       errorLog,
		infoLog:        infoLog,
		models:         models,
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	// Configure http server
	// Routing is configured with httprouter in the routes.go file
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     errorLog,
	}

	// Start HTTP server
	infoLog.Printf("starting %s on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	infoLog.Fatal(err)
}
