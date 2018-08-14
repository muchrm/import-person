package personsql

import (
	"database/sql"
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
	history_education string
	history_work      string
}

func GetPersonInfo(db *sql.DB) {
	var person Person
	rows, err := db.Query(
		`SELECT
			Person.personId,
			Person.personCode,
			Person.fName,
			Person.lName,
			PersonT.fName2,
			PersonT.lName2,
			PersonT.birthDate,
			rs_academic_position.acp_name_th,
			PersonT.emailAddr AS email,
			(
				CASE Person.fStatus
				WHEN 1 THEN 'ยังปฏิบัติงานอยู่'
				WHEN 2 THEN 'ไม่ได้ปฏิบัติงานแล้ว'
				ELSE null
				END
			) AS PersonStatus,
			CONCAT(
			'[',
			GROUP_CONCAT(
			JSON_OBJECT(
			'levelName',Level.levelName,
			'levelNameEng',Level.levelNameEng,
			'degreeName',Degree.degreeName,
			'educmajorName',Educmajor.educmajorName,
			'educplaceName',Educplace.educplaceName,
			'countryName',Country.countryName
			)),']') AS history_education,
			CONCAT(
			'[',
			GROUP_CONCAT(
			JSON_OBJECT(
			'start_date',moresearcher.rs_experience.ex_start_date,
			'end_date',moresearcher.rs_experience.ex_end_date,
			'position',moresearcher.rs_experience.ex_position,
			'workplace',moresearcher.rs_experience.ex_workplace
			)),']') AS history_work
		FROM Person
		LEFT JOIN PersonT ON PersonT.personId = Person.personId
		LEFT JOIN rs_academic ON rs_academic.ac_ps_id = Person.personId
		LEFT JOIN rs_academic_position ON rs_academic_position.acp_id = rs_academic.ac_acp_id
		LEFT JOIN Education ON Education.personId = Person.personId
		LEFT JOIN Level ON Education.levelId = Level.levelId
		LEFT JOIN Degree ON Education.degreeId = Degree.degreeId
		LEFT JOIN Educmajor ON Education.educmajorId = Educmajor.educmajorId
		LEFT JOIN Educplace ON Education.educplaceId = Educplace.educplaceId
		LEFT JOIN Country ON Education.countryId = Country.countryId
		LEFT JOIN moresearcher.rs_user ON moresearcher.rs_user.us_ps_id = Person.personId
		LEFT JOIN moresearcher.rs_experience ON moresearcher.rs_experience.ex_us_id = moresearcher.rs_user.us_id
		GROUP BY Person.personId`)
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