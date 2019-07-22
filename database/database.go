package database

import (
	"../php2go"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func Connect(name string) *sql.DB {
	link := get(name)
	dbDsn := dsn(link)
	db, err := sql.Open(link["type"], dbDsn)
	if err != nil {
		panic(err)
	}
	php2go.Dump(db)
	return db
}

func Mysql() *sql.DB {
	return Connect("mysql")
}

func Redis() *sql.DB {
	return Connect("redis")
}
