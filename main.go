package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/people")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var (
		id           int
		prefixName   sql.NullString
		adminposName sql.NullString
		deptName     sql.NullString
	)
	rows, err := db.Query(`SELECT
	Person.personId,
	Prefix.prefixName,
	Adminpos.adminposName,
	Department.deptName
FROM
	Person
	Left Join Prefix ON Person.prefixId = Prefix.prefixId
	Left Join Adminpos ON Person.adminposId = Adminpos.adminposId
	Left Join Department ON Person.deptId = Department.deptId
WHERE Person.personId > 0
ORDER BY
	Person.adminId ASC,
	Person.levelposId DESC,
	Person.salary DESC`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &prefixName, &adminposName, &deptName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, prefixName, adminposName, deptName)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
