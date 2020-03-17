# x-oss

---

> 依赖
```
需要go 1.14
```

> 测试
```
go run x-oss
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
缩放：resize=20%,20% (width,height)支持百分比（urlencode=%25）及固定参数,height可省略
模糊：blur=2 (distance)模糊程度，数值越大越模糊
裁剪：thumb=500,500,1200,1200 (x1,y1,x2,y2) 坐标以左上角为(0,0)
翻转：flip=0,1 (x,y) 支持0 or 1，1表示该轴翻转
字符画：ascii=1 (isdo)
```
