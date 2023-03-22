package data

import (
	"database/sql"
	"fmt"
)

type VotesModel struct {
	DB *sql.DB
}

func (v VotesModel) Insert(vote *Vote) error {
	query := `
	INSERT INTO votes (session_id, product_id, vote) 
	VALUES ($1,$2,$3)`

	args := []interface{}{vote.SessionID, vote.ProductID, vote.Vote}

	_, err := v.DB.Exec(query, args...)
	if err != nil {
		err = fmt.Errorf("error inserting vote in db: %w", err)
	}
	return err
}
