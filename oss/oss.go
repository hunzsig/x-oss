package oss

import (
	"../php2go"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

/**
 * 分析文件，并返回文件信息
 */
func AnalysisFile(file multipart.File, header *multipart.FileHeader) (map[string]string, error) {
	fileInfo := make(map[string]string)
	if file == nil {
		newFile, err := header.Open()
		if err != nil {
			return fileInfo, err
		}
		file = newFile
		defer file.Close()
	}
	php2go.Dump(header.Header)
	fileNameSep := php2go.Explode(".", header.Filename)
	// 后缀名
	fileInfo["suffix"] = fileNameSep[len(fileNameSep)-1]
	fileNameSep = fileNameSep[:len(fileNameSep)-1]
	// 文件名
	fileInfo["name"] = php2go.Implode(".", fileNameSep)
	fileInfo["token_name"] = strconv.FormatInt(php2go.Microtime(), 10)
	// 文件大小
	fileInfo["size"] = strconv.FormatInt(header.Size, 10)
	fileSha1, err := php2go.Sha1FileSrc(file)
	fileInfo["sha1"] = fileSha1
	sha1Arr := php2go.Split(fileSha1, 4)
	// 文件路径
	fileInfo["path"] = "./uploads/" + php2go.Implode("/", sha1Arr) + "/"
	fileInfo["uri"] = fileInfo["path"] + fileInfo["token_name"] + "." + fileInfo["suffix"]
	php2go.Dump(fileInfo)
	out, err := os.OpenFile(fileInfo["uri"], os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fileInfo, err
	}
	defer out.Close()
	io.Copy(out, file)
	return fileInfo, nil
}
