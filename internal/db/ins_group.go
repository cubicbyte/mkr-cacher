package db

// InsertGroup inserts the group to the DB
func (db *DB) InsertGroup(structureId int, facultyId int, course int, id int, name string, priority int, educationForm int) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Stmt(db.stmtGroup).Exec(structureId, facultyId, course, id, name, priority, educationForm); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
