package scope

import (
	"../database"
	"../php2go"
	"../response"
	"github.com/kataras/iris"
	"io"
	"mime/multipart"
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
 * 分析文件，并返回文件信息
 */
func analysisFile(file multipart.File, header multipart.FileHeader) (map[string]string, error) {
	fileInfo := make(map[string]string)
	if file == nil {
		newFile, err := header.Open()
		if err != nil {
			return fileInfo, err
		}
		file = newFile
	}
	fileNameSep := php2go.Explode(".", header.Filename)
	// 后缀名
	fileInfo["suffix"] = fileNameSep[len(fileNameSep)-1]
	fileNameSep = fileNameSep[:len(fileNameSep)-1]
	// 文件名
	fileInfo["name"] = php2go.Implode(".", fileNameSep)
	// 文件大小
	fileInfo["size"] = string(header.Size)
	fileSha1, err := php2go.Sha1FileSrc(file)
	fileInfo["sha1"] = fileSha1
	php2go.Dump(fileSha1)
	sha1Arr := php2go.Split(fileSha1, 4)
	php2go.Dump(sha1Arr)
	// 文件路径
	fileInfo["path"] = "./uploads/" + php2go.Implode("/", sha1Arr) + "/"
	out, err := os.OpenFile(fileInfo["path"]+fileInfo["name"], os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fileInfo, err
	}
	defer out.Close()
	io.Copy(out, file)
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
	fileNameSep := php2go.Explode(".", header.Filename)
	// 后缀名
	fileSuffix := fileNameSep[len(fileNameSep)-1]
	fileNameSep = fileNameSep[:len(fileNameSep)-1]
	// 文件名
	fileName := php2go.Implode(".", fileNameSep)
	// 文件大小
	fileSize := header.Size
	fileContent, err := php2go.Sha1FileSrc(file)
	php2go.Dump(fileContent)
	sha1Arr := php2go.Split(fileContent, 4)
	php2go.Dump(sha1Arr)
	// 文件路径
	filePath := "./uploads/" + php2go.Implode("/", sha1Arr) + "/"

	out, err := os.OpenFile(filePath+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	defer out.Close()
	io.Copy(out, file)
	return response.Success(ctx, string(fileSize), nil)
}

/**
 * 上传文件（多个）
 */
func UploadMulti(ctx iris.Context) bool {
	// Get the file from the request.
	file, header, err := ctx.FormFile("file")
	php2go.Dump(file)
	panic("3")
	if err != nil {
		return response.Error(ctx, err.Error(), nil)
	}
	defer file.Close()
	fileNameSep := php2go.Explode(".", header.Filename)
	// 后缀名
	fileSuffix := fileNameSep[len(fileNameSep)-1]
	fileNameSep = fileNameSep[:len(fileNameSep)-1]
	// 文件名
	fileName := php2go.Implode(".", fileNameSep)
	// 文件大小
	fileSize := header.Size
	fileContent, err := php2go.Sha1FileSrc(file)
	php2go.Dump(fileContent)
	sha1Arr := php2go.Split(fileContent, 4)
	php2go.Dump(sha1Arr)
	// 文件路径
	filePath := "./uploads/" + php2go.Implode("/", sha1Arr) + "/"

	out, err := os.OpenFile(filePath+fileName, os.O_WRONLY|os.O_CREATE, 0666)
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
