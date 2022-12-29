package api

type Structure struct {
	Id        int    `json:"id"`
	ShortName string `json:"shortName"`
	FullName  string `json:"fullName"`
}

type Faculty struct {
	Structure
}

type Course struct {
	Course int `json:"course"`
}

type Group struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Course        int    `json:"course"`
	Priority      int    `json:"priority"`
	EducationForm int    `json:"educationForm"`
}

type Chair struct {
	Structure
}

type Person struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	LastName   string `json:"lastName"`
}

type Student struct {
	Person
}

type CallSchedule struct {
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
	Number    int    `json:"number"`
	Length    int    `json:"length"`
}

type TimeTableDate struct {
	Date    string            `json:"date"`
	Lessons []TimeTableLesson `json:"lessons"`
}

type TimeTableLesson struct {
	Description string            `json:"description"`
	Number      int               `json:"number"`
	Periods     []TimeTablePeriod `json:"periods"`
}

type TimeTablePeriod struct {
	R1                    int    `json:"r1"`
	Rz14                  int    `json:"rz14"`
	Rz15                  int    `json:"rz15"`
	R5                    int    `json:"r5"`
	DisciplineId          int    `json:"disciplineId"`
	EducationDisciplineId int    `json:"educationDisciplineId"`
	DisciplineFullName    string `json:"disciplineFullName"`
	DisciplineShortName   string `json:"disciplineShortName"`
	Classroom             string `json:"classroom"`
	TimeStart             string `json:"timeStart"`
	TimeEnd               string `json:"timeEnd"`
	TeachersName          string `json:"teachersName"`
	TeachersNameFull      string `json:"teachersNameFull"`
	Type                  int    `json:"type"`
	DateUpdated           string `json:"dateUpdated"`
	NonstandartTime       bool   `json:"nonstandartTime"`
	Groups                string `json:"groups"`
	ExtraText             bool   `json:"extraText"`
}

type Classroom struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CountPlace int    `json:"countPlace"`
	Type       int    `json:"type"`
}

type Rd struct {
	Html string `json:"html"`
}

type TeacherByName struct {
	ChairName  string `json:"chairName"`
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	LastName   string `json:"lastName"`
}