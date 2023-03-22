package data

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Sessions interface {
		Insert(session *Session) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Sessions: SessionModel{
			DB: db,
		},
	}
}

type SessionModel struct {
	DB *sql.DB
}

func (s SessionModel) Insert(session *Session) error {
	query := `
	INSERT INTO sessions (session_id) 
	VALUES ($1)`

	args := []interface{}{session.ID}

	result, err := s.DB.Exec(query, args...)
	fmt.Println(result)

	return err
}
func (s SessionModel) Get(id int64) (*Session, error) {

	return nil, nil
}
