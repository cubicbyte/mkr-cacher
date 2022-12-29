package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Url string
}

func (api *Api) makeRequest(method string, path string, body io.Reader, result any) error {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "uk")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := res.Body.Close(); err != nil {
		return err
	}

	if err := json.Unmarshal(resBody, &result); err != nil {
		return err
	}

	return nil
}

func (api *Api) ListStructures() (*[]Structure, error) {
	var result []Structure

	err := api.makeRequest("GET", api.Url+"/list/structures", nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) ListFaculties(structureId int) (*[]Faculty, error) {
	var jsonParams = []byte(fmt.Sprint(`{"structureId":`, structureId, "}"))
	var result []Faculty

	err := api.makeRequest("POST", api.Url+"/list/faculties", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) ListCourses(facultyId int) (*[]Course, error) {
	var jsonParams = []byte(fmt.Sprint(`{"facultyId":`, facultyId, "}"))
	var result []Course

	err := api.makeRequest("POST", api.Url+"/list/courses", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) ListGroups(facultyId int, course int) (*[]Group, error) {
	var jsonParams = []byte(fmt.Sprint(`{"facultyId":`, facultyId, `,"course":`, course, "}"))
	var result []Group

	err := api.makeRequest("POST", api.Url+"/list/groups", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) ListChairs() (*[]Chair, error) {
	var result []Chair

	err := api.makeRequest("POST", api.Url+"/list/chairs", nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) ListTeachersByChair(chairId int) (*[]Person, error) {
	var jsonParams = []byte(fmt.Sprint(`{"chairId":`, chairId, "}"))
	var result []Person

	err := api.makeRequest("POST", api.Url+"/list/teachers-by-chair", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) ListStudentsByGroup(groupId int) (*[]Student, error) {
	var jsonParams = []byte(fmt.Sprint(`{"groupId":`, groupId, "}"))
	var result []Student

	err := api.makeRequest("POST", api.Url+"/list/students-by-group", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) TimeTableCallSchedule() (*[]CallSchedule, error) {
	var result []CallSchedule

	err := api.makeRequest("POST", api.Url+"/time-table/call-schedule", nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) TimeTableGroup(groupId int, dateStart string, dateEnd string) (*[]TimeTableDate, error) {
	var jsonParams = []byte(fmt.Sprint(`{"groupId":`, groupId, `,"dateStart":"`, dateStart, `","dateEnd":"`, dateEnd, "\"}"))
	var result []TimeTableDate

	err := api.makeRequest("POST", api.Url+"/time-table/group", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) TimeTableStudent(studentId int, dateStart string, dateEnd string) (*[]TimeTableDate, error) {
	var jsonParams = []byte(fmt.Sprint(`{"studentId":`, studentId, `,"dateStart":"`, dateStart, `","dateEnd":"`, dateEnd, "\"}"))
	var result []TimeTableDate

	err := api.makeRequest("POST", api.Url+"/time-table/student", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) TimeTableTeacher(teacherId int, dateStart string, dateEnd string) (*[]TimeTableDate, error) {
	var jsonParams = []byte(fmt.Sprint(`{"teacherId":`, teacherId, `,"dateStart":"`, dateStart, `","dateEnd":"`, dateEnd, "\"}"))
	var result []TimeTableDate

	err := api.makeRequest("POST", api.Url+"/time-table/teacher", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) TimeTableClassroom(classroomId int, dateStart string, dateEnd string) (*[]TimeTableDate, error) {
	var jsonParams = []byte(fmt.Sprint(`{"classroomId":`, classroomId, `,"dateStart":"`, dateStart, `","dateEnd":"`, dateEnd, "\"}"))
	var result []TimeTableDate

	err := api.makeRequest("POST", api.Url+"/time-table/classroom", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) TimeTableFreeClassroom(structureId int, corpusId int, lessonNumberStart int, lessonNumberEnd int, date string) (*[]Classroom, error) {
	var jsonParams = []byte(fmt.Sprint(`{"structureId":`, structureId, `,"corpusId":`, corpusId, `,"lessonNumberStart":`, lessonNumberStart, `,"lessonNumberEnd":`, lessonNumberEnd, `,"date":"`, date, "\"}"))
	var result []Classroom

	err := api.makeRequest("POST", api.Url+"/time-table/free-classroom", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) TimeTableScheduleAd(r1 int, r2 int) (*Rd, error) {
	var jsonParams = []byte(fmt.Sprint(`{"r1":`, r1, `,"r2":`, r2, "}"))
	var result Rd

	err := api.makeRequest("POST", api.Url+"/time-table/schedule-ad", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *Api) OtherSearchTeachers(name string) (*[]TeacherByName, error) {
	var jsonParams = []byte(fmt.Sprint(`{"name":"`, name, "\"}"))
	var result []TeacherByName

	err := api.makeRequest("POST", api.Url+"/other/search-teachers", bytes.NewBuffer(jsonParams), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
