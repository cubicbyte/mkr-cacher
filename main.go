package main

import (
	"encoding/json"
	"fmt"
	"os"
	"schedule-cacher/internal/db"
	"schedule-cacher/pkg/api"
)

// If api is down, here is another one https://api.kname.edu.ua
const ApiUrl = "https://mia.mobil.knute.edu.ua"
const DbFile = "test.db"

func main() {
	myApi := api.Api{Url: ApiUrl}

	db, err := db.NewDB(DbFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//const GroupId = 7157
	//const Date = "2022-12-27"
	res, err := myApi.ListGroups(7, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//jsonRes, _ := json.MarshalIndent(*res, "", "  ")
	jsonRes, _ := json.Marshal(*res)
	fmt.Println(string(jsonRes))

	for _, group := range *res {
		if err := db.InsertGroup(0, 7, 2, group.Id, group.Name, group.Priority, group.EducationForm); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if err := db.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
