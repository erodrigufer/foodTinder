package data

import (
	"database/sql"
	"errors"
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

func (v VotesModel) Votes(sessionID string) ([]Vote, error) {
	query := `
	SELECT session_id, product_id, vote
	FROM votes
	WHERE session_id = $1`

	var vote Vote
	votes := make([]Vote, 0, 25)
	rows, err := v.DB.Query(query, sessionID)
	if err != nil {
		return votes, fmt.Errorf("error performing db query: %w", err)
	}

	for rows.Next() {
		err := rows.Scan(&vote.SessionID, &vote.ProductID, &vote.Vote)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return votes, nil
			default:
				return votes, fmt.Errorf("error performing scan of rows: %w", err)
			}
		}
		votes = append(votes, vote)
	}

	return votes, nil
}
