package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// writeJSON, encode data to JSON, send back response with status code as
// defined in parameter.
func writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	// Encode data to JSON.
	js, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error while encoding data to json: %w", err)
	}
	// JSON-specific HTTP header.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
