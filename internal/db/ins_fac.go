package database

import (
	"github.com/cubicbyte/mkr-cacher/pkg/api"
)

// InsertFaculty inserts the faculty to the DB
func (db *DB) InsertFaculty(structureId int, f *api.Faculty) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Stmt(db.stmtFac).Exec(structureId, f.Id, f.ShortName, f.FullName); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
