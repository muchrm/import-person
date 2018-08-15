package personsql

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/muchrm/import-person/person"
)

type PersonSQL struct {
	personCode       string
	fName            string
	lName            string
	fName2           string
	lName2           string
	birthDate        string
	prefixName       sql.NullString
	email            string
	PersonStatus     string
	historyEducation sql.NullString
	historyWork      sql.NullString
}

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/people")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Query(db *sql.DB, path string) (*sql.Rows, error) {
	bufQuery, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	query := string(bufQuery)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func GetPersonInfo() ([]person.Person, error) {
	var personSQL PersonSQL
	var err error
	persons := []person.Person{}

	db, err := GetConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	rows, err := Query(db, "./query/person.sql")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&personSQL.personCode,
			&personSQL.personCode,
			&personSQL.fName,
			&personSQL.lName,
			&personSQL.fName2,
			&personSQL.lName2,
			&personSQL.birthDate,
			&personSQL.prefixName,
			&personSQL.email,
			&personSQL.PersonStatus,
			&personSQL.historyEducation,
			&personSQL.historyWork,
		)
		if err != nil {
			return nil, err
		}
		log.Println(personSQL)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return persons, err
}
