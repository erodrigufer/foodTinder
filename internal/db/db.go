package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// OpenDB, creates a connection pool to a db.
// dsn stands for data source name of db to which we want to establish a
// connection.
func OpenDB(dsn string) (*sql.DB, error) {
	// Create connection pool to db.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool to db: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Verify connection to db.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in ping to db: %w", err)
	}

	return db, nil
}
