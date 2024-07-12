package database

import "database/sql"

type Database interface {
	Init(user, password, host, dbname string) error
	GetDB() *sql.DB
}
