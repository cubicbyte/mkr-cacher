package database

const setup_query = `
	DROP TABLE IF EXISTS Schedule;
	DROP TABLE IF EXISTS Structures;
	DROP TABLE IF EXISTS Faculties;
	DROP TABLE IF EXISTS Courses;
	DROP TABLE IF EXISTS Groups;

	CREATE TABLE Schedule (
		groupId int,
		date date,
		lessons json,
		PRIMARY KEY (groupId, date)
	);

	CREATE TABLE Structures (
	    id int,
	    shortName varchar(16),
	    fullName varchar(32),
	    PRIMARY KEY (id)
	);

	CREATE TABLE Faculties (
	    structureId int,
	    id int,
	    shortName varchar(16),
	    fullName varchar(32),
	    PRIMARY KEY (id)
	);

	CREATE TABLE Courses (
	    facultyId int,
	    course int,
	    PRIMARY KEY (facultyId, course)
	);

	CREATE TABLE Groups (
		facultyId int,
		id int,
		name varchar(16),
		course int,
		priority int,
		educationForm int,
		PRIMARY KEY (id)
	);

	CREATE INDEX IF NOT EXISTS schedule_groupId ON Schedule(groupId);
	CREATE INDEX IF NOT EXISTS schedule_date 	ON Schedule(date);
	CREATE INDEX IF NOT EXISTS structure_id 	ON Faculties(structureId);
	CREATE INDEX IF NOT EXISTS faculty_id 		ON Courses(facultyId);
	CREATE INDEX IF NOT EXISTS faculty_id 		ON Groups(facultyId);
	CREATE INDEX IF NOT EXISTS course 			ON Groups(course);`

const ins_sched_query = `
	INSERT OR IGNORE INTO Schedule (
		groupId,
		date,
		lessons
	) VALUES (
		?, ?, ?
	);`

const ins_struct_query = `
	INSERT INTO Structures (
		id,
	    shortName,
	    fullName
	) VALUES (
		?, ?, ?
	);`

const ins_fac_query = `
	INSERT INTO Faculties (
		structureId,
		id,
	    shortName,
	    fullName
	) VALUES (
		?, ?, ?, ?
	);`

const ins_course_query = `
	INSERT INTO Courses (
		facultyId,
	    course
	) VALUES (
		?, ?
	);`

const ins_group_query = `
	INSERT OR IGNORE INTO Groups (
		facultyId,
		id,
		name,
		course,
		priority,
		educationForm
	) VALUES (
		?, ?, ?, ?, ?, ?
	);`
