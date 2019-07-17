package scope

import (
	"../database"
	"../php2go"
	"../response"
	"github.com/kataras/iris"
	"io"
	"os"
)

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
	out, err := os.OpenFile("./uploads/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
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
	php2go.VarDump(result)
	// token := ctx.Params().Get("token")
	return response.Download(ctx, "./uploads/test.txt")
}
