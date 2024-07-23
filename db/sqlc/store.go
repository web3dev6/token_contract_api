package db

import (
	"database/sql"
)

// a generic interfce for store
type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries and transactions - a real db (postgres in app)
type SQLStore struct {
	*Queries // extend struct functionality in golang - inheritance equivalent
	db       *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{db: db, Queries: New(db)}
}
