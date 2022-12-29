package db

func (db *DB) Close() error {
	defer func() {
		db.stmtSched.Close()
		db.stmtGroup.Close()
		db.sql.Close()
	}()

	return nil
}
