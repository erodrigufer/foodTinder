package api

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"

	"github.com/erodrigufer/foodTinder/internal/data"
)

// charsetSession, valid character-set for generating random session IDs.
const charsetSession = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// newRandomString, returns a random string.
// Parameters: length is the number of characters of the string that should be
// returned. charset, is the valid character set from which to generate a
// random string.
func newRandomString(length int, charset string) (string, error) {
	// Make a slice of length length, in which to store random characters.
	b := make([]byte, length)
	for i := range b {
		// Use the cryptographically more secure implementation rand.Int() to
		// get a pseudo-random integer (this is more secure than seeding a
		// pseudo-random generator yourself).
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("error: getting a random integer failed: %w", err)
		}
		// Convert the big.Int type to a int64 type.
		r64 := r.Int64()
		// Pick a single character from the character-set through indexing the
		// string of the character-set. The index is a random number between 0
		// and the length of the character-set minus 1.
		b[i] = charset[int(r64)]
	}

	return string(b), nil
}

// newRandomSession, returns a random session ID of a given character length,
// using solely characters defined in a charset.
func newRandomSession(length int, charset string) (string, error) {
	s, err := newRandomString(length, charset)
	if err != nil {
		return "", fmt.Errorf("error: could not create random session: %w", err)
	}
	return s, nil
}

// createNewSession, create a new unique session upon request. Store the new
// session persistently in a database, send session back to client.
func (app *Application) createNewSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := newRandomSession(64, charsetSession)
	if err != nil {
		// Create fail response.
		resp := data.SessionResponse{
			APIVersion: "1.0",
			Status:     "fail",
		}
		app.ErrorLog.Printf("error creating new session: %v", err)
		// Respond with Internal Server Error status code.
		err = writeJSON(w, http.StatusInternalServerError, resp)
		return
	}
	// Create successful response.
	resp := data.SessionResponse{
		APIVersion: "1.0",
		Status:     "success",
		Data: data.SessionResponseData{
			SessionID: sessionID,
		},
	}

	err = writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.serverError(w, err)
		return
	}

}
