package data

import (
	"database/sql"
	"errors"
	"fmt"
)

type SessionModel struct {
	DB *sql.DB
}

func (s SessionModel) Insert(session *Session) error {
	query := `
	INSERT INTO sessions (session_id) 
	VALUES ($1)`

	args := []interface{}{session.ID}

	_, err := s.DB.Exec(query, args...)
	if err != nil {
		err = fmt.Errorf("error inserting session in db: %w", err)
	}
	return err
}

func (s SessionModel) Exists(sessionID string) (bool, error) {
	query := `
	SELECT * 
	FROM sessions
	WHERE session_id = $1`

	// These two variables are only required to properly run the Scan method.
	var id int64
	var sid string
	err := s.DB.QueryRow(query, sessionID).Scan(&id, &sid)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return false, ErrRecordNotFound
		default:
			return false, err
		}
	}

	return true, nil
}
