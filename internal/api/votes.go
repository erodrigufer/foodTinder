package api

import (
	"fmt"
	"net/http"

	"github.com/erodrigufer/foodTinder/internal/data"
	"github.com/julienschmidt/httprouter"
)

// createNewVote, create a new unique vote upon request. Store the new
// vote persistently in a database.
func (app *Application) createNewVote(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Session_ID string `json:"session_id"`
		Vote       bool   `json:"vote"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		errorDisplayClient := fmt.Sprintf("%s\nError: %s\n", http.StatusText(http.StatusBadRequest), err)
		http.Error(w, errorDisplayClient, http.StatusBadRequest)
		return
	}

	idParam := httprouter.ParamsFromContext(r.Context())
	productID := idParam.ByName("id")

	vote := &data.Vote{
		SessionID: input.Session_ID,
		Vote:      input.Vote,
		ProductID: productID,
	}

	err = app.models.Votes.Insert(vote)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Create successful response.
	resp := data.SessionResponse{
		APIVersion: API_VERSION,
		Status:     "success",
	}

	err = writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.serverError(w, err)
		return
	}

}
