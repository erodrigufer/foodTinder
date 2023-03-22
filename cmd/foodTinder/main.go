package main

import (
	"log"
	"os"

	"github.com/erodrigufer/foodTinder/internal/api"
)

func main() {

	// Create a logger for INFO messages, the prefix "INFO" and a tab will be
	// displayed before each log message. The flags Ldate and Ltime provide the
	// local date and time.
	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create an ERROR messages logger, additionally use the Lshortfile flag to
	// display the file's name and line number for the error.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	srv := api.NewServer(8000)

	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
