package api

import (
	"fmt"
	"net/http"
	"time"
)

// NewServer, return a new instance of the server.
// port, defines the port at which the server will listen.
func NewServer(port int) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv
}
