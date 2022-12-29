package db

// InsertSchedule inserts the schedule for the day into DB
func (db *DB) InsertSchedule(groupId int, date string, lessons string) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Stmt(db.stmtSched).Exec(groupId, date, lessons); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
