# x-oss

---

> 安装依赖
```
go get -u github.com/kataras/iris
go get -u github.com/go-sql-driver/mysql
go get -u github.com/lib/pq
go get -u github.com/denisenkom/go-mssqldb # need golang.org/x/crypto/md4
go get -u github.com/mattn/go-sqlite3 # need gcc++
go get -u github.com/gomodule/redigo/redis
go get -u github.com/jinzhu/gorm
```

> mac 版本不对应
(如compile: version "go1.12" does not match go tool version "go1.13")，
可以敲句
```
export GOROOT=/usr/local/opt/go/libexec
```

> 构建/执行项目
```
windos
# go build
# x-oss.exe

unix / mac
# go build
# ./x-oss
```