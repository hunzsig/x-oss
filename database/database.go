package database

import (
	"database/sql"
	"fmt"
)

func dsn(conf map[string]string){
	dbDsn := fmt.Sprintf("port=32768 user=%s password=%s dbname=%s sslmode=disable", "postgres", "pass123", "postgres")
	db, err := sql.Open(conf["type"], dbDsn)
}