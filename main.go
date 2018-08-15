package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muchrm/import-person/personsql"
)

func main() {
	result, err := personsql.GetPersonInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
