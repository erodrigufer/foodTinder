package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes, define the routing of the server.
func routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/session", createNewSession)

	return router
}
