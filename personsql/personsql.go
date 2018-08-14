package personsql

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetPersons() {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/people")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	GetPersonInfo(db)
}

type Person struct {
	personId          int
	personCode        string
	fName             string
	lName             string
	fName2            string
	lName2            string
	birthDate         string
	acp_name_th       sql.NullString
	email             string
	PersonStatus      string
	history_education sql.NullString
	history_work      sql.NullString
}

func GetPersonInfo(db *sql.DB) {
	var person Person
	bufQuery, err := ioutil.ReadFile("./query/person.sql")
	if err != nil {
		log.Fatal(err)
	}
	query := string(bufQuery)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&person.personCode,
			&person.personCode,
			&person.fName,
			&person.lName,
			&person.fName2,
			&person.lName2,
			&person.birthDate,
			&person.acp_name_th,
			&person.email,
			&person.PersonStatus,
			&person.history_education,
			&person.history_work,
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(
			person.personCode,
			person.fName,
			person.lName,
			person.fName2,
			person.lName2,
			person.birthDate,
			person.acp_name_th,
			person.email,
			person.PersonStatus,
			person.history_education,
			person.history_work,
		)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
