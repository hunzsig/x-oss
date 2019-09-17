package scope

import (
	"../database"
	"../oss"
	"../php2go"
	"../response"
	"github.com/kataras/iris"
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
 * 上传文件（一个）
 */
func UploadOne(ctx iris.Context) bool {
	// Get the file from the request.
	file, header, err := ctx.FormFile("file")
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	defer file.Close()
	fileInfo, err := oss.AnalysisFile(file, header)
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	return response.Success(ctx, fileInfo.Size, nil)
}

/**
 * 上传文件（多个）
 */
func UploadMulti(ctx iris.Context) bool {
	//获取通过iris.WithPostMaxMemory获取的最大上传值大小。
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	/*
		form := ctx.Request().MultipartForm
		files := form.File["files[]"]
		failures := 0
		for _, file := range files {
			_, err = saveUploadedFile(file, "./uploads")
			if err != nil {
				failures++
				ctx.Writef("failed to upload: %s\n", file.Filename)
			}
		}
	*/
	return true
}

/**
 * 根据token下载文件
 */
func Download(ctx iris.Context, fileKey string) bool {
	result := database.Mysql().Connect.Table("files").Row()
	php2go.Dump(result)
	// token := ctx.Params().Get("token")
	return response.Download(ctx, "./uploads/test.txt")
}
