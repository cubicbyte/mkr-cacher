package database

import "github.com/cubicbyte/mkr-cacher/pkg/api"

// InsertCourse inserts the course to the DB
func (db *DB) InsertCourse(facultyId int, c *api.Course) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Stmt(db.stmtCourse).Exec(facultyId, c.Course); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
