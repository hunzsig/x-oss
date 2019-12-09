# x-oss

---

> 安装依赖
```
go get -u github.com/kataras/iris
go get -u github.com/go-sql-driver/mysql
go get -u github.com/lib/pq
go get -u github.com/denisenkom/go-mssqldb # need golang.org/x/crypto/md4
go get -u github.com/mattn/go-sqlite3   // 需要 gcc++ http://www.mingw.org
go get -u github.com/gomodule/redigo/redis
go get -u github.com/jinzhu/gorm
go get -u github.com/hunzsig/graphics
```

```
报golang.org/x/net错请：
git clone https://github.com/golang/net.git golang.org/x/net(备用https://github.com/hunzsig/net.git)

报golang.org/x/crypto错请：
git clone https://github.com/golang/crypto.git golang.org/x/crypto(备用https://github.com/hunzsig/crypto.git)
其他类似~~
其中golang.org/路径是gopath对应存放的目录，根据本地目录自行调整结构
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

> 路由参数
```
上传文件（一个）：/oss/<USER_TOKEN>/upload/one
上传文件（多个）：/oss/<USER_TOKEN>/upload/multi
下载文件（一个）：/oss/<USER_TOKEN>/download/<FILE_TOKEN>
```

> 下载图片参数
```
反相：reverse=1 (isdo)
灰度：grayscale=1 (isdo)
缩放：resize=20%,20% (width,height)支持百分比及固定参数,height可省略
模糊：blur=2 (distance)模糊程度，数值越大越模糊
裁剪：thumb=500,500,1200,1200 (x1,y1,x2,y2) 坐标以左上角为(0,0)
翻转：flip=0,1 (x,y) 支持0 or 1，1表示该轴翻转
字符画：ascii=1 (isdo)
```
