package sql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muchrm/import-person/sql/info"
)

func GetPerson() {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/people")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	info.GetPersonInfo(db)
}
