package database

import (
	"database/sql"
)

type DB struct {
	sql        *sql.DB
	stmtSched  *sql.Stmt
	stmtStruct *sql.Stmt
	stmtFac    *sql.Stmt
	stmtCourse *sql.Stmt
	stmtGroup  *sql.Stmt
}

type FullGroup struct {
	FacultyId     int    `json:"facultyId"`
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Course        int    `json:"course"`
	Priority      int    `json:"priority"`
	EducationForm int    `json:"educationForm"`
}
