package database

import (
	"../system"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(link string) {
	currentConf := conf(link)
	dbDsn := dsn(currentConf)
	db, err := sql.Open(currentConf["type"], dbDsn)
	if err != nil {
		panic(err)
	}
	system.Dump(db)
}
