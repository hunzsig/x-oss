package scope

import (
	"../database"
	"../models"
	"../oss"
	"../php2go"
	"../response"
	"github.com/kataras/iris"
	"os"
	"time"
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
	fileInfo, err := oss.AnalysisFile(file, header)
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	// record db
	fileInfo.UserToken = ctx.Params().Get("user_token")
	fileInfo.FromUrl = ""
	fileInfo.CallQty = "0"
	fileInfo.CallLastTime = time.Now().Format("2006-01-02 15:04:05")
	fileInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	fileInfo.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	database.Mysql().Connect.Save(&fileInfo)
	defer database.Mysql().Connect.Close()
	returnData := make(map[string]string)
	returnData["key"] = fileInfo.Key
	returnData["size"] = fileInfo.Size
	returnData["name"] = fileInfo.Name
	return response.Success(ctx, "upload_ok", returnData)
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
	files := models.Files{}
	database.Mysql().Connect.Where("`key` = ?", fileKey).First(&files)
	if files.Hash == "" {
		response.NotFound(ctx, "resource not found", nil)
		return false
	}
	if files.Uri == "" {
		response.NotFound(ctx, "resource has a bad uri", nil)
		return false
	}
	if oss.IsExist(files.Uri) == false {
		response.NotFound(ctx, "resource not exist", nil)
		return false
	}
	return response.Download(ctx, files.Uri)
}
