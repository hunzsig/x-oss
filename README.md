# x-oss

---

> 安装依赖
```
go get -u github.com/kataras/iris
go get -u github.com/go-sql-driver/mysql
go get -u github.com/lib/pq
go get -u github.com/denisenkom/go-mssqldb # need golang.org/x/crypto/md4
go get -u github.com/mattn/go-sqlite3   // 需要 gcc++
go get -u github.com/gomodule/redigo/redis
go get -u github.com/jinzhu/gorm
go get -u github.com/hunzsig/graphics
```

> mac 版本不对应
(如compile: version "go1.12" does not match go tool version "go1.13")，
这是由于你的go版本与工具不对应造成的

> 如果是IDE内部终端报错，去「设置」查看GOROOT是否选择了正确的GO sdk

> 如果是全局环境不对可以敲句
```
export GOROOT=/usr/local/opt/go/libexec
```

> cannot find package "golang.org/x/crypto/md4" 问题,由于golang.org被强，无法直接go get获得
```
进入你的GOPATH
git clone https://github.com/golang/crypto.git golang.org/x/crypto
其中golang.org/x/crypto是对应存放的目录，根据本地目录自行调整结构
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