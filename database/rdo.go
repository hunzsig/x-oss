package database

import (
	"../mapping"
	"github.com/gomodule/redigo/redis"
)

// DB 对象
type rdo struct {
	dbType  string
	dsn     string
	options map[string]string
	obj     redis.Conn
}

// 连接,并返回一个GDO
func RDO(name string) *rdo {
	link := get(name)
	dbDsn := dsn(link)
	if link["type"] != mapping.DBType.Redis.Value {
		panic("not support db type: " + link["type"])
	}
	conn, err := redis.Dial("tcp", "10.1.210.69:6379")
	if err != nil {
		panic("connect redis error :" + err.Error())
	}
	defer conn.Close()
	rdo := new(rdo)
	rdo.dbType = link["type"]
	rdo.dsn = dbDsn
	rdo.obj = conn
	return rdo
}

func Redis() *rdo {
	return RDO("redis")
}

func Cache() *rdo {
	return RDO("cache")
}
