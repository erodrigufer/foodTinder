package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes, define the routing of the server.
func (app *Application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/session", app.createNewSession)

	return router
}
