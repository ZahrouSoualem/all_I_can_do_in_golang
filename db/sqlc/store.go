package db

import "database/sql"

type Store interface {
	Querier
}

type DBStore struct {
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &DBStore{
		Queries: New(db),
	}
}
