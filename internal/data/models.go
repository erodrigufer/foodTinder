package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Sessions interface {
		Insert(session *Session) error
		Exists(sessionID string) (bool, error)
	}
	Votes interface {
		Insert(vote *Vote) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Sessions: SessionModel{
			DB: db,
		},
		Votes: VotesModel{
			DB: db,
		},
	}
}
