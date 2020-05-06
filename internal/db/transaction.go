package db

import (
	"database/sql"
)

type Transaction interface {
}

type txImpl struct {
	tx *sql.Tx
}
