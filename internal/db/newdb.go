package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
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
	// Read all sql query files
	setupSql, err := readFileCont(filepath.Join("sql", "setup.sql"))
	if err != nil {
		return nil, err
	}
	schedSql, err := readFileCont(filepath.Join("sql", "ins_sched.sql"))
	if err != nil {
		return nil, err
	}
	groupSql, err := readFileCont(filepath.Join("sql", "ins_group.sql"))
	if err != nil {
		return nil, err
	}

	// Open DB file
	sqlDB, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	// Exec setup query to setup DB
	if _, err = sqlDB.Exec(*setupSql); err != nil {
		return nil, err
	}

	// Prepare all other queries
	stmtSched, err := sqlDB.Prepare(*schedSql)
	if err != nil {
		return nil, err
	}
	stmtGroup, err := sqlDB.Prepare(*groupSql)
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
