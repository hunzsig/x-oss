package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"../system"
)

func connect(link string) {
	currentConf := conf(link)
	dbDsn := dsn(currentConf)
	db, err := sql.Open(currentConf["type"], dbDsn)
	system.Dump(db)
}
