package db

import (
	"cmd/internal/storage/dbs"
	"context"
	"database/sql"
	"log"
)

type Repository struct {
	connection *sql.DB
	queries    *dbs.Queries
}

func NewRepository(connection *sql.DB) (*Repository, error) {
	var queries, err = dbs.Prepare(context.Background(), connection)
	if err != nil {
		return nil, err
	}

	return &Repository{connection, queries}, nil
}

func MustRepository(connection *sql.DB) *Repository {
	var repository, err = NewRepository(connection)
	if err != nil {
		panic(err)
	}
	return repository
}

func (r *Repository) Connection() *sql.DB {
	return r.connection
}

func (r *Repository) Queries() *dbs.Queries {
	return r.queries
}

func (r *Repository) Close() error {
	return r.queries.Close()
}

func (r *Repository) WithTransaction(ctx context.Context, fn func(queries *dbs.Queries) error) error {
	return withTransaction(ctx, r.connection, r.queries, fn)
}

func withTransaction(ctx context.Context, db *sql.DB, queries *dbs.Queries, fn func(queries *dbs.Queries) error) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			err := tx.Rollback()
			if err != nil {
				log.Println("Rollback failed: ", err)
				return
			}

			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			err := tx.Rollback()
			if err != nil {
				log.Println("Rollback failed: ", err)
				return
			}
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(queries.WithTx(tx))

	return err
}
