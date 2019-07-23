package mapping

type dbType struct {
	Mysql  mapping
	Pgsql  mapping
	Sqlite mapping
	Mssql  mapping
	Mongo  mapping
	Redis  mapping
}

var (
	DBType dbType
)

func init() {
	// db type
	DBType = dbType{
		Mysql: mapping{
			Value: "mysql",
			Label: "mysql",
		},
		Pgsql: mapping{
			Value: "pgsql",
			Label: "pgsql",
		},
		Sqlite: mapping{
			Value: "sqlite",
			Label: "sqlite",
		},
		Mssql: mapping{
			Value: "mssql",
			Label: "mssql",
		},
		Redis: mapping{
			Value: "redis",
			Label: "redis",
		},
	}
}
