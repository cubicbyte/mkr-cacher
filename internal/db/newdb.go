package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// Returns the string content of the file
func readFileCont(file string) (*string, error) {
	raw, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cont := string(raw)
	return &cont, nil
}

// NewDB is DB struct "constructor"
func NewDB(dbFile string) (*DB, error) {

	// Open DB file
	sqlDB, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	// Exec setup query to setup DB
	if _, err = sqlDB.Exec(setup_query); err != nil {
		return nil, err
	}

	// Prepare all other queries
	stmtSched, err := sqlDB.Prepare(ins_sched_query)
	if err != nil {
		return nil, err
	}
	stmtGroup, err := sqlDB.Prepare(ins_group_query)
	if err != nil {
		return nil, err
	}

	// Return DB instance
	db := DB{
		sql:       sqlDB,
		stmtSched: stmtSched,
		stmtGroup: stmtGroup,
	}

	return &db, nil
}
