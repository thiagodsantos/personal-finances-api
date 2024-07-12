package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

func Init(user, password, host, dbname string) *sql.DB {
	if db == nil {
		dsn := user + ":" + password + "@tcp(" + host + ")/" + dbname
		var err error
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}

func GetDB() *sql.DB {
	return db
}
