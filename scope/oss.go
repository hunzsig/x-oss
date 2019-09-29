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
	"strconv"
	"strings"
	"time"
)

const tempImagesRoot = "./my_data/temp_images"

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
func Download(ctx iris.Context) bool {
	fileKey := ctx.Params().Get("fileKey")
	fileInfo := models.Files{}
	database.Mysql().Connect.Where("`key` = ?", fileKey).First(&fileInfo)
	if fileInfo.Hash == "" {
		response.NotFound(ctx, "resource not found", nil)
		return false
	}
	if fileInfo.Uri == "" {
		response.NotFound(ctx, "resource has a bad uri", nil)
		return false
	}
	if oss.IsExist(fileInfo.Uri) == false {
		response.NotFound(ctx, "resource not exist", nil)
		return false
	}

	// 更新调用
	callQTY, _ := strconv.Atoi(fileInfo.CallQty)
	fileInfo.CallQty = strconv.Itoa(callQTY + 1)
	database.Mysql().Connect.Table("files").Where("`hash` = ?", fileInfo.Hash).Updates(models.Files{
		CallQty:      strconv.Itoa(callQTY + 1),
		CallLastTime: time.Now().Format("2006-01-02 15:04:05"),
	})

	fmt.Println(strings.Index(fileInfo.ContentType, "image"))

	// 图片处理
	if strings.Index(fileInfo.ContentType, "image") >= 0 {

		// 检测form values
		var imagesChange []string

		// 灰度
		colorGrayscale := ctx.FormValue("cg")
		if colorGrayscale == "1" {
			imagesChange = append(imagesChange, "cg")
		}
		// 反转
		colorReverse := ctx.FormValue("cr")
		if colorReverse == "1" {
			imagesChange = append(imagesChange, "cr")
		}

		imagesChangeStr := "_" + php2go.Implode("_", imagesChange)
		tempUri := tempImagesRoot + "/" + fileInfo.Hash + imagesChangeStr + "." + fileInfo.Suffix
		err := os.MkdirAll(tempImagesRoot, os.ModePerm)
		if err != nil {
			response.Error(ctx, err.Error(), nil)
			return false
		}
		f, _ := os.Create(tempUri)
		defer f.Close()
		err = oss.ImageEncode(tempUri, f, oss.ImageColorReverse(fileInfo.Uri))
		if err != nil {
			response.Error(ctx, err.Error(), nil)
			return false
		}
		return response.Download(ctx, tempUri)
	}
	php2go.Dump(fileInfo)
	// 原库中文件
	return response.Download(ctx, fileInfo.Uri)
}
