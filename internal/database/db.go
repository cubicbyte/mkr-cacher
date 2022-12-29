package database

import (
	"database/sql"
)

type DB struct {
	sql       *sql.DB
	stmtSched *sql.Stmt
	stmtGroup *sql.Stmt
}

type FullGroup struct {
	StructureId   int    `json:"structureId"`
	FacultyId     int    `json:"facultyId"`
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Course        int    `json:"course"`
	Priority      int    `json:"priority"`
	EducationForm int    `json:"educationForm"`
}
