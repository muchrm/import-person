package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muchrm/import-person/personmongo"
	"github.com/muchrm/import-person/personsql"
)

func main() {
	persons, err := personsql.GetPersonInfo()
	if err != nil {
		log.Fatal(err)
	}
	for _, person := range persons {
		err = personmongo.AddPerson(person)
		if err != nil {
			log.Fatal(err)
		}
	}
}
