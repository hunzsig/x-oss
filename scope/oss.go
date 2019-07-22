package scope

import (
	"../database"
	"../php2go"
	"../response"
	"github.com/kataras/iris"
	"io"
	"mime/multipart"
	"os"
	"strconv"
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
 * 分析文件，并返回文件信息
 */
func analysisFile(file multipart.File, header *multipart.FileHeader) (map[string]string, error) {
	fileInfo := make(map[string]string)
	if file == nil {
		newFile, err := header.Open()
		if err != nil {
			return fileInfo, err
		}
		file = newFile
		defer file.Close()
	}
	fileNameSep := php2go.Explode(".", header.Filename)
	// 后缀名
	fileInfo["suffix"] = fileNameSep[len(fileNameSep)-1]
	fileNameSep = fileNameSep[:len(fileNameSep)-1]
	// 文件名
	fileInfo["name"] = php2go.Implode(".", fileNameSep)
	fileInfo["token_name"] = strconv.FormatInt(php2go.Microtime(), 10)
	// 文件大小
	fileInfo["size"] = string(header.Size)
	fileSha1, err := php2go.Sha1FileSrc(file)
	fileInfo["sha1"] = fileSha1
	sha1Arr := php2go.Split(fileSha1, 4)
	// 文件路径
	fileInfo["path"] = "./uploads/" + php2go.Implode("/", sha1Arr) + "/"
	fileInfo["uri"] = fileInfo["path"] + fileInfo["token_name"] + "." + fileInfo["suffix"]
	out, err := os.OpenFile(fileInfo["uri"], os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fileInfo, err
	}
	defer out.Close()
	io.Copy(out, file)
	php2go.Dump(fileInfo)
	return fileInfo, nil
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
	fileInfo, err := analysisFile(file, header)
	return response.Success(ctx, fileInfo["size"], nil)
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
	result, err := database.Mysql().Query("select * from `test`")
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	php2go.Dump(result)
	// token := ctx.Params().Get("token")
	return response.Download(ctx, "./uploads/test.txt")
}
