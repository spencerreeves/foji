// Code generated by foji 0.4, template: foji/pgx/db.go.tpl; DO NOT EDIT.

package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// DB is the common interface for database operations.
type DB interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

func IsErrNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows) || errors.Is(err, sql.ErrNoRows)
}

type Repo struct {
	db DB
}

func New(db DB) Repo {
	return Repo{db: db}
}
