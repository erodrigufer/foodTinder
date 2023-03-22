package api

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// newServer, return a new instance of the server.
// port, defines the port at which the server will listen.
func newServer(port int, errLog *log.Logger) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ErrorLog:     errLog,
		Handler:      routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv
}

// serverError, sends an error message and stack trace to the error logger and
// then sends a generic 500 Internal Server Error response to the client.
func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// The first parameter of Output equals the calldepth, which is the count
	// of the number of frames to skip when computing the file name
	// and line number. So basically, just go back on the stack trace to display
	// the name of function (file) which called the error logging helper
	// function.
	app.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
