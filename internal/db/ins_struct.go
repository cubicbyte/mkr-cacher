package database

import (
	"github.com/cubicbyte/mkr-cacher/pkg/api"
)

// InsertStructure inserts the structure to the DB
func (db *DB) InsertStructure(st *api.Structure) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Stmt(db.stmtStruct).Exec(st.Id, st.ShortName, st.FullName); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
