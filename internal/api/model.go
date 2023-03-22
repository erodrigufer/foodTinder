package api

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/erodrigufer/foodTinder/internal/data"
	"github.com/erodrigufer/foodTinder/internal/db"
	_ "github.com/lib/pq"
)

const API_VERSION = "1.0"

// Application, object to interact with API from within main cmd file.
type Application struct {
	// Srv, *http.Server
	Srv *http.Server
	// ErrorLog, log server errors.
	ErrorLog *log.Logger
	// InfoLog, informative server logger.
	InfoLog *log.Logger
	// DB, DB connection pool.
	DB *sql.DB

	models data.Models
}

// NewApplication, create a new Application struct that hosts the loggers and
// http.Server
func NewApplication(port int, dsn string) *Application {
	app := new(Application)
	// Create a logger for INFO messages, the prefix "INFO" and a tab will be
	// displayed before each log message. The flags Ldate and Ltime provide the
	// local date and time.
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create an ERROR messages logger, additionally use the Lshortfile flag to
	// display the file's name and line number for the error.
	app.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.Srv = app.newServer(port)

	db_conn, err := db.OpenDB(dsn)
	if err != nil {
		app.ErrorLog.Fatalf("unable to establish db connection: %v", err)
	}
	app.DB = db_conn

	app.InfoLog.Printf("Database connection pool successfully established.")

	app.models = data.NewModels(app.DB)

	return app
}
