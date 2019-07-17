package scope

import (
	"../database"
	"../php2go"
	"../response"
	"github.com/kataras/iris"
	"io"
	"os"
)

func init() {
	if php2go.IsDir("./uploads") == false {
		err := php2go.Mkdir("uploads", os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
}

/**
 * 上传文件
 */
func Upload(ctx iris.Context) bool {
	// Get the file from the request.
	file, header, err := ctx.FormFile("file")
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	defer file.Close()
	filename := header.Filename
	fileSize := header.Size
	fileContent, err := php2go.Sha1FileSrc(file)
	php2go.Dump(fileContent)
	sha1Arr := php2go.Split(fileContent, 4)
	php2go.Dump(sha1Arr)
	path := php2go.Implode("/", sha1Arr)
	out, err := os.OpenFile("./uploads/"+path+"/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	defer out.Close()
	io.Copy(out, file)
	return response.Success(ctx, string(fileSize), nil)
}

/**
 * 根据token下载文件
 */
func Download(ctx iris.Context) bool {
	result, err := database.Mysql().Query("select * from `test`")
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	php2go.Dump(result)
	// token := ctx.Params().Get("token")
	return response.Download(ctx, "./uploads/test.txt")
}
