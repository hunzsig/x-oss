package scope

import (
	"../database"
	"../models"
	"../oss"
	"../php2go"
	"../response"
	"github.com/kataras/iris"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const tempImagesRoot = "./my_data/temp_images"

/**
 * 清理temp文件
 */
func clearTemp(tempDate string) {
	rd, _ := ioutil.ReadDir(tempImagesRoot)
	for _, fi := range rd {
		if fi.IsDir() && fi.Name() != tempDate {
			_ = os.RemoveAll(tempImagesRoot + "/" + fi.Name())
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
		thumb := ctx.FormValue("thumb")
		resize := ctx.FormValue("resize")
		rotate := ctx.FormValue("rotate")
		blur := ctx.FormValue("blur")
		colorGrayscale := ctx.FormValue("grayscale")
		colorReverse := ctx.FormValue("reverse")
		ascii := ctx.FormValue("ascii")

		var thumbX1 int
		var thumbX2 int
		var thumbY1 int
		var thumbY2 int
		var resizeIx int
		var resizeIy int
		var resizeFx float64
		var resizeFy float64
		var rotateAngle int
		var blurDistance float64

		thumbX1 = 0
		thumbX2 = 0
		thumbY1 = 0
		thumbY2 = 0
		thumbI := false

		resizeIx = 0
		resizeIy = 0
		resizeFx = 0.00
		resizeFy = 0.00

		resizeI := false
		resizeF := false

		blurDistance = 0.0

		// 裁剪
		if thumb != "" {
			thumbSplit := php2go.Explode(",", thumb)
			if len(thumbSplit) != 4 {
				response.Error(ctx, "thumb has 4 vector", nil)
				return false
			}
			for idx, v := range thumbSplit {
				if v != "" {
					thumbSplit[idx] = php2go.Trim(v, " ")
				}
			}
			thumbX1, _ = strconv.Atoi(thumbSplit[0])
			thumbY1, _ = strconv.Atoi(thumbSplit[1])
			thumbX2, _ = strconv.Atoi(thumbSplit[2])
			thumbY2, _ = strconv.Atoi(thumbSplit[3])
			imagesChange = append(imagesChange, "tb", strconv.Itoa(thumbX1), strconv.Itoa(thumbY1), strconv.Itoa(thumbX2), strconv.Itoa(thumbY2))
			thumbI = true
		}
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
					"rs",
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
						"rs",
						strconv.FormatFloat(resizeFx, 'f', 2, 64),
						strconv.FormatFloat(resizeFy, 'f', 2, 64),
					)
					resizeF = true
				} else {
					resizeIx, _ = strconv.Atoi(resizeSplit[0])
					resizeIy, _ = strconv.Atoi(resizeSplit[1])
					imagesChange = append(imagesChange, "rs", strconv.Itoa(resizeIx), strconv.Itoa(resizeIy))
					resizeI = true
				}
			}
		}
		// 旋转
		if rotate != "" {
			rotateAngle, _ = strconv.Atoi(rotate)
			imagesChange = append(imagesChange, "ro")
			imagesChange = append(imagesChange, strconv.Itoa(rotateAngle))
			fileInfo.Suffix = "png" // 旋转时固定png，获得透明背景
		}
		// 模糊
		if blur != "" {
			blurDistance, _ = strconv.ParseFloat(blur, 1)
			imagesChange = append(imagesChange, "bl")
			imagesChange = append(imagesChange, strconv.FormatFloat(blurDistance, 'f', 1, 64))
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
			imagesChange = append(imagesChange, "ii")
			fileInfo.Suffix = "txt" // ascii时固定txt，获得文本
		}

		// 无改动时，返回原图
		if len(imagesChange) == 0 {
			return response.Download(ctx, fileInfo.Uri, fileInfo.Name+"."+fileInfo.Suffix)
		}

		// 构建临时目录,临时文件保留一个月，过后清理
		imagesChangeStr := "_" + php2go.Implode("_", imagesChange)
		tempDate := time.Now().Format("200601")

		// clear temp
		go clearTemp(tempDate)

		tempPath := tempImagesRoot + "/" + tempDate + "/"
		if oss.IsExist(tempPath) == false {
			err := os.MkdirAll(tempPath, os.ModePerm)
			if err != nil {
				response.Error(ctx, err.Error(), nil)
				return false
			}
		}
		tempUri := tempPath + fileInfo.Hash + imagesChangeStr + "." + fileInfo.Suffix
		if oss.IsExist(tempUri) == true {
			return response.Download(ctx, tempUri, fileInfo.Name+"."+fileInfo.Suffix)
		}

		rgba := oss.ImageRGBA(fileInfo.Uri)

		// 裁剪
		if thumbI {
			rgba = oss.ImageThumb(rgba, thumbX1, thumbY1, thumbX2, thumbY2)
		}
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
		// 模糊
		if blurDistance != 0.0 {
			rgba = oss.ImageBlur(rgba, blurDistance)
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
			err := oss.ImageAscii(f, rgba)
			if err != nil {
				response.Error(ctx, err.Error(), nil)
				return false
			}
			// images
		} else {
			err := oss.ImageEncode(f, rgba, fileInfo.Suffix)
			if err != nil {
				response.Error(ctx, err.Error(), nil)
				return false
			}
		}
		return response.Download(ctx, tempUri, fileInfo.Name+"."+fileInfo.Suffix)
	}
	// 原库中文件
	return response.Download(ctx, fileInfo.Uri, fileInfo.Name+"."+fileInfo.Suffix)
}
