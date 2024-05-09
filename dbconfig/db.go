package dbconfig

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	host = os.Getenv("HOST")
	password = os.Getenv("PASSWORD")
	dbname = os.Getenv("DBNAME")
)

func DB() (*sql.DB, error) {
	
	connectionString := fmt.Sprintf("host=%s user=postgres password=%s " + "dbname=%s sslmode=disable", host, password, dbname)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return db, err
}