package scope

import (
	"../database"
	"../models"
	"../oss"
	"../php2go"
	"../response"
	"fmt"
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
	fileInfo, err := oss.AnalysisFile(ctx, file, header)
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}

	// return info
	returnData := make(map[string]string)
	returnData["key"] = fileInfo.Key
	returnData["size"] = fileInfo.Size
	returnData["name"] = fileInfo.Name
	returnData["suffix"] = fileInfo.Suffix
	return response.Success(ctx, "upload_ok", returnData)
}

/**
 * 上传文件（多个）
 */
func UploadMulti(ctx iris.Context) bool {
	//获取通过iris.WithPostMaxMemory获取的最大上传值大小。
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	fmt.Println(maxSize)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	form := ctx.Request().MultipartForm
	files := form.File["files[]"]
	success := 0
	failures := 0
	var returnData []map[string]string
	for _, file := range files {
		fileInfo, err := oss.AnalysisFile(ctx, nil, file)
		if err != nil {
			failures++
			returnData = append(returnData, nil)
		} else {
			success++
			item := make(map[string]string)
			item["key"] = fileInfo.Key
			item["size"] = fileInfo.Size
			item["name"] = fileInfo.Name
			item["suffix"] = fileInfo.Suffix
			returnData = append(returnData, item)
		}
	}
	return response.Success(ctx, "upload_ok", returnData)
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
