package userprofile

import (
	"database/sql"
	"fmt"
)

type BasicUserProfile struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func ConnectDB() {

	fmt.Println("Connecting to Mysql Database")

	db, error := sql.Open("mysql", "root:ueducation@tcp(localhost:3306)/producer")

	if error != nil {
		panic(error.Error())
	}

	defer db.Close()
}

func SelectQuery() {
	fmt.Println("Connecting to Mysql Database for select query")

	db, error := sql.Open("mysql", "root:ueducation@tcp(localhost:3306)/producer")

	if error != nil {
		panic(error.Error())
	}

	results, error := db.Query("select * from basic_user_profile")

	if error != nil {
		panic(error.Error())
	}

	for results.Next() {
		var basicUserProfile BasicUserProfile

		err := results.Scan(&basicUserProfile.ID, &basicUserProfile.Email, &basicUserProfile.Name, &basicUserProfile.PhoneNumber)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(basicUserProfile)
	}
}

func SingleRow() {
	fmt.Println("Connectind to DB for fetching single row")

	db, error := sql.Open("mysql", "root:ueducation@tcp(localhost:3306)/producer")

	if error != nil {
		panic(error.Error())
	}

	var basicUserProfile BasicUserProfile
	db.QueryRow("select * from basic_user_profile where id = ?", 10).Scan(&basicUserProfile.ID, &basicUserProfile.Email, &basicUserProfile.Name, &basicUserProfile.PhoneNumber)

	fmt.Println(basicUserProfile)
}
