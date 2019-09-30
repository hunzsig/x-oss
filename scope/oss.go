package scope

import (
	"../database"
	"../models"
	"../oss"
	"../php2go"
	"../response"
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

	// 图片处理
	if strings.Index(fileInfo.ContentType, "image") >= 0 {

		// 检测form values
		var imagesChange []string

		// get value
		resize := ctx.FormValue("resize")
		rotate := ctx.FormValue("rotate")
		colorGrayscale := ctx.FormValue("grayscale")
		colorReverse := ctx.FormValue("reverse")
		ascii := ctx.FormValue("ascii")

		// 缩放
		var resizeIx int
		var resizeIy int
		var resizeFx float64
		var resizeFy float64
		var rotateAngle int

		resizeIx = 0
		resizeIy = 0
		resizeFx = 0.00
		resizeFy = 0.00

		resizeI := false
		resizeF := false

		// 缩放
		if resize != "" {
			resizeSplit := php2go.Explode(",", resize)
			if len(resizeSplit) == 1 {
				resizeSplit = append(resizeSplit, "0")
			}
			for idx, v := range resizeSplit {
				if v != "" {
					resizeSplit[idx] = php2go.Trim(v, " ")
				}
			}

			if php2go.Strpos(resizeSplit[0], "%", 0) > -1 { // 百分比形式时
				resizeSplit[0] = php2go.StrReplace("%", "", resizeSplit[0], -1)
				resizeSplit[1] = php2go.StrReplace("%", "", resizeSplit[1], -1)
				resizeFx, _ = strconv.ParseFloat(resizeSplit[0], 2)
				resizeFy, _ = strconv.ParseFloat(resizeSplit[1], 2)
				resizeFx = resizeFx * 0.01
				resizeFy = resizeFy * 0.01
				imagesChange = append(
					imagesChange,
					"cg",
					strconv.FormatFloat(resizeFx, 'f', 2, 64),
					strconv.FormatFloat(resizeFy, 'f', 2, 64),
				)
				resizeF = true
			} else {
				resizeFx, _ = strconv.ParseFloat(resizeSplit[0], 2)
				resizeFy, _ = strconv.ParseFloat(resizeSplit[1], 2)
				if resizeFx < 1.00 {
					imagesChange = append(
						imagesChange,
						"cg",
						strconv.FormatFloat(resizeFx, 'f', 2, 64),
						strconv.FormatFloat(resizeFy, 'f', 2, 64),
					)
					resizeF = true
				} else {
					resizeIx, _ = strconv.Atoi(resizeSplit[0])
					resizeIy, _ = strconv.Atoi(resizeSplit[1])
					imagesChange = append(imagesChange, "cg", strconv.Itoa(resizeIx), strconv.Itoa(resizeIy))
					resizeI = true
				}
			}
		}
		// 旋转
		if rotate != "" {
			rotateAngle, _ = strconv.Atoi(rotate)
			imagesChange = append(imagesChange, "ro")
		}
		// 灰度
		if colorGrayscale == "1" {
			imagesChange = append(imagesChange, "cg")
		}
		// 色相反转
		if colorReverse == "1" {
			imagesChange = append(imagesChange, "cr")
		}
		// ascii
		if ascii == "1" {
			imagesChange = append(imagesChange, "ascii")
		}

		// 无改动时，返回原图
		if len(imagesChange) == 0 {
			return response.Download(ctx, fileInfo.Uri)
		}

		// 构建临时目录
		imagesChangeStr := "_" + php2go.Implode("_", imagesChange)
		err := os.MkdirAll(tempImagesRoot, os.ModePerm)
		if err != nil {
			response.Error(ctx, err.Error(), nil)
			return false
		}
		if ascii == "1" {
			fileInfo.Suffix = "txt"
		}
		tempUri := tempImagesRoot + "/" + fileInfo.Hash + imagesChangeStr + "." + fileInfo.Suffix
		if oss.IsExist(tempUri) == true {
			return response.Download(ctx, tempUri)
		}

		rgba := oss.ImageRGBA(fileInfo.Uri)

		// 缩放
		if resizeI {
			rgba = oss.ImageResizeInt(rgba, resizeIx, resizeIy)
		} else if resizeF {
			rgba = oss.ImageResizeFloat(rgba, resizeFx, resizeFy)
		}
		// 旋转
		if rotateAngle != 0 {
			rgba = oss.ImageRotate(rgba, rotateAngle)
		}
		// 灰度
		if colorGrayscale == "1" {
			rgba = oss.ImageColorGrayscale(rgba)
		}
		// 色相反转
		if colorReverse == "1" {
			rgba = oss.ImageColorReverse(rgba)
		}

		f, _ := os.Create(tempUri)
		defer f.Close()

		// ascii
		if ascii == "1" {
			err = oss.ImageAscii(f, rgba)
			if err != nil {
				response.Error(ctx, err.Error(), nil)
				return false
			}
			// images
		} else {
			err = oss.ImageEncode(tempUri, f, rgba)
			if err != nil {
				response.Error(ctx, err.Error(), nil)
				return false
			}
		}
		return response.Download(ctx, tempUri)
	}
	php2go.Dump(fileInfo)
	// 原库中文件
	return response.Download(ctx, fileInfo.Uri)
}
