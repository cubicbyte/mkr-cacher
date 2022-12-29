package db

import (
	"database/sql"
)

type DB struct {
	sql       *sql.DB
	stmtSched *sql.Stmt
	stmtGroup *sql.Stmt
}
