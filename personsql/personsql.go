package personsql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/muchrm/import-person/person"
)

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
			&personSQL.position,
			&personSQL.email,
			&personSQL.PersonStatus,
			&personSQL.historyEducation,
			&personSQL.historyWork,
		)
		if err != nil {
			return nil, err
		}
		teacher, err := parsePerson(personSQL)
		if err == nil {
			persons = append(persons, teacher)
		}
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
	if personSQL.position.Valid {
		teacher.OfficerPosition = personSQL.position.String
	}
	teacher.Email = personSQL.email
	splitedEmail := strings.Split(teacher.Email, "@")
	if len(splitedEmail) > 0 {
		teacher.OfficerLogin = splitedEmail[0]
	}
	teacher.OfficerStatus = personSQL.PersonStatus

	teacher.HistoryEducations, err = parseHistoryEducations(personSQL)
	if err != nil {
		return person.Person{}, err
	}
	teacher.HistoryWorks, err = parseHistoryWork(personSQL)
	if err != nil {
		return person.Person{}, err
	}
	return teacher, nil
}
func parseHistoryEducations(personSQL PersonSQL) ([]person.HistoryEducation, error) {
	if !personSQL.historyEducation.Valid {
		return []person.HistoryEducation{}, nil
	}
	var historyEducationSQLs []HistoryEducationSQL
	historyEducations := []person.HistoryEducation{}
	err := json.Unmarshal([]byte(personSQL.historyEducation.String), &historyEducationSQLs)
	if err != nil {
		return []person.HistoryEducation{}, err
	}
	for _, historyEducationSQL := range historyEducationSQLs {
		historyEducation := person.HistoryEducation{}
		historyEducation.LevelName = historyEducationSQL.LevelName
		historyEducation.DegreeName = historyEducationSQL.DegreeName
		historyEducation.CountryName = historyEducationSQL.CountryName
		historyEducation.MajorName = historyEducationSQL.MajorName
		historyEducation.PlaceName = historyEducationSQL.PlaceName
		endYear, err := strconv.Atoi(historyEducationSQL.EndYear)
		if err == nil {
			historyEducation.EndYear = endYear
		}
		historyEducations = append(historyEducations, historyEducation)
	}
	return historyEducations, nil
}
func parseHistoryWork(personSQL PersonSQL) ([]person.HistoryWork, error) {
	if !personSQL.historyWork.Valid {
		return []person.HistoryWork{}, nil
	}
	historyWorks := []person.HistoryWork{}
	var historyWorkSQLs []HistoryWorkSQL
	err := json.Unmarshal([]byte(personSQL.historyWork.String), &historyWorkSQLs)
	if err != nil {
		return []person.HistoryWork{}, err
	}

	for _, historyWorkSQL := range historyWorkSQLs {
		historyWork := person.HistoryWork{}
		historyWork.Position = historyWorkSQL.Position
		historyWork.Workplace = historyWorkSQL.Workplace
		if startDate, err := converDate(historyWorkSQL.StartDate); err == nil {
			historyWork.StartDate = startDate
		}
		if endDate, err := converDate(historyWorkSQL.EndDate); err == nil {
			historyWork.EndDate = endDate
		}
		if historyWorkSQL.EndDate == "0000-00-00" {
			historyWork.DateLess = true
		} else {
			historyWork.DateLess = false
		}
		historyWorks = append(historyWorks, historyWork)
	}
	return historyWorks, nil
}
func converDate(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", s))
}
