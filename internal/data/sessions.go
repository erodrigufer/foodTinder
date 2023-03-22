package data

import (
	"database/sql"
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
func (s SessionModel) Get(id int64) (*Session, error) {

	return nil, nil
}
