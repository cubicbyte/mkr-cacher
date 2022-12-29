package database

import "fmt"

// InsertGroups inserts the groups to the DB
func (db *DB) InsertGroups(groups *[]FullGroup) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	for _, g := range *groups {
		if _, err := tx.Stmt(db.stmtGroup).Exec(g.StructureId, g.FacultyId, g.Id, g.Name, g.Course, g.Priority, g.EducationForm); err != nil {
			tx.Rollback()
			fmt.Println(g)
			return err
		}
	}

	return tx.Commit()
}
