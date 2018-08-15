package personsql

import (
	"database/sql"
	"io/ioutil"
	"strconv"

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

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/people")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func query(db *sql.DB, path string) (*sql.Rows, error) {
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

	db, err := getConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	rows, err := query(db, "./query/person.sql")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
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
		teacher, err := parsePerson(personSQL)
		persons = append(persons, teacher)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return persons, err
}
func parsePerson(personSQL PersonSQL) (person.Person, error) {
	teacher := person.Person{}
	personCode, err := strconv.Atoi(personSQL.personCode)
	if err != nil {
		return person.Person{}, err
	}
	teacher.OfficerCode = personCode
	teacher.OfficerName = personSQL.fName
	teacher.OfficerSurname = personSQL.lName
	teacher.OfficerNameEng = personSQL.fName2
	teacher.OfficerSurnameEng = personSQL.lName2
	if personSQL.prefixName.Valid {
		teacher.OfficerPrefixName = personSQL.prefixName.String
	}
	teacher.Email = personSQL.email
	teacher.OfficerStatus = personSQL.PersonStatus

	return teacher, nil
}
