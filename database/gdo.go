package database

import (
	"../mapping"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

// DB 对象
type gdo struct {
	DbType  string
	Dsn     string
	Connect *gorm.DB
	Options map[string]string
}

// 连接,并返回一个GDO
func GDO(name string) *gdo {
	link := get(name)
	dbDsn := dsn(link)
	gormType := ""
	switch link["type"] {
	case mapping.DBType.Mysql.Value:
		gormType = "mysql"
	case mapping.DBType.Pgsql.Value:
		gormType = "postgres"
	case mapping.DBType.Mssql.Value:
		gormType = "mssql"
	case mapping.DBType.Sqlite.Value:
		gormType = "sqlite3"
	default:
		panic("not support db type: " + link["type"])
	}
	con, err := gorm.Open(gormType, dbDsn)
	if err != nil {
		panic(err)
	}
	// Disable table name's pluralization
	con.SingularTable(true)

	// link poor
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	con.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	con.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	con.DB().SetConnMaxLifetime(time.Hour)

	gdo := new(gdo)
	gdo.DbType = link["type"]
	gdo.Dsn = dbDsn
	gdo.Connect = con
	gdo.Options = make(map[string]string)
	return gdo
}

// ---------------------------------------------

/**
 * db type
 */
func Mysql() *gdo {
	return GDO("mysql")
}

func Pgsql() *gdo {
	return GDO("pgsql")
}

func Mssql() *gdo {
	return GDO("mssql")
}

func Sqlite() *gdo {
	return GDO("sqlite")
}
