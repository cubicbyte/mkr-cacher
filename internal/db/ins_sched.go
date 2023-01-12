package database

import (
	"encoding/json"
	"github.com/cubicbyte/mkr-cacher/pkg/api"
)

// InsertSchedule inserts the schedule for the day into DB
func (db *DB) InsertSchedule(groupId int, date string, lessons *[]api.TimeTableLesson) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	jsonRes, _ := json.Marshal(*lessons)
	lessonsJson := string(jsonRes)

	if _, err := tx.Stmt(db.stmtSched).Exec(groupId, date, lessonsJson); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
