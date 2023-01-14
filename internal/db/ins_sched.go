package database

import (
	"encoding/json"
	"github.com/cubicbyte/mkr-cacher/pkg/api"
)

// InsertSchedule inserts the schedule for the day into DB
func (db *DB) InsertSchedule(groups *[]FullGroup, api *api.Api, dateStart string, dateEnd string) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	for _, g := range *groups {
		sched, err := api.TimeTableGroup(g.Id, dateStart, dateEnd)
		if err != nil {
			tx.Rollback()
			return err
		}

		for _, s := range *sched {
			bytesJson, _ := json.Marshal(s.Lessons)
			lessonsJson := string(bytesJson)

			if _, err := tx.Stmt(db.stmtSched).Exec(g.Id, s.Date, lessonsJson); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}
