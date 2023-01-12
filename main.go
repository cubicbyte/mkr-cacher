package main

import (
	"fmt"
	"github.com/cubicbyte/mkr-cacher/internal/db"
	"github.com/cubicbyte/mkr-cacher/pkg/api"
	"log"
	"os"
)

// If api is down, here is another one https://api.kname.edu.ua
const ApiUrl = "https://mia.mobil.knute.edu.ua"
const DbFile = "test.db"

// Group buffer. Groups are added to it when * function is called.
// When filled, the groups are transferred to the database.
var groupBuff = make([]database.FullGroup, 0, 1024)

func insertGroup(db *database.DB, sId int, fId int, g *api.Group) error {
	var fGroup = database.FullGroup{
		StructureId:   sId,
		FacultyId:     fId,
		Id:            g.Id,
		Name:          g.Name,
		Course:        g.Course,
		Priority:      g.Priority,
		EducationForm: g.EducationForm,
	}

	groupBuff = append(groupBuff, fGroup)
	if len(groupBuff) == cap(groupBuff) {
		err := db.InsertGroups(&groupBuff)
		groupBuff = groupBuff[:0]
		return err
	}

	return nil
}

func main() {
	myApi := api.Api{Url: ApiUrl}

	db, err := database.NewDB(DbFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println("Loading groups...")
	structures, err := myApi.ListStructures()
	if err != nil {
		log.Panicf("Can't get structures: %s\n", err)
	}

	for _, structure := range *structures {
		faculties, err := myApi.ListFaculties(structure.Id)
		if err != nil {
			log.Panicf("Can't get faculties: %s\n", err)
		}

		for _, faculty := range *faculties {
			courses, err := myApi.ListCourses(faculty.Id)
			if err != nil {
				log.Panicf("Can't get courses: %s\n", err)
			}

			for _, course := range *courses {
				groups, err := myApi.ListGroups(faculty.Id, course.Course)
				if err != nil {
					log.Panicf("Can't get groups: %s\n", err)
				}

				for _, group := range *groups {
					if err := insertGroup(db, structure.Id, faculty.Id, &group); err != nil {
						log.Fatalf("Can't insert group: %s\n", err)
					}
				}
			}
		}
	}

	err = db.InsertGroups(&groupBuff)
	if err != nil {
		log.Fatalf("Can't insert groups: %s\n", err)
	}
	groupBuff = groupBuff[:0]
	log.Println("Groups loading completed!")

	if err := db.Close(); err != nil {
		log.Panicf("Can't close DB connection: %s\n", err)
	}
}
