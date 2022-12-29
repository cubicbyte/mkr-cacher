CREATE TABLE IF NOT EXISTS Schedule (
    groupId int NOT NULL,
    date date NOT NULL,
    lessons json,
    PRIMARY KEY (groupId, date)
);

CREATE TABLE IF NOT EXISTS Groups (
    structureId int NOT NULL,
    facultyId int NOT NULL,
    id int NOT NULL,
    name varchar(32) NOT NULL,
    course int NOT NULL,
    priority int NOT NULL,
    educationForm int NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS schedule_groupId ON Schedule(groupId);
CREATE INDEX IF NOT EXISTS schedule_date ON Schedule(date);

CREATE INDEX IF NOT EXISTS group_structureId ON Groups(structureId);
CREATE INDEX IF NOT EXISTS group_facultyId ON Groups(facultyId);
CREATE INDEX IF NOT EXISTS group_course ON Groups(course);
