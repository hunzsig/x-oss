package oss

import (
	"../models"
	"../php2go"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

/**
 * 分析文件，并返回文件信息
 */
func AnalysisFile(file multipart.File, header *multipart.FileHeader) (models.Files, error) {
	php2go.Dump("AnalysisFile")
	fileInfo := models.Files{}
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
	fileInfo.Suffix = fileNameSep[len(fileNameSep)-1]
	fileNameSep = fileNameSep[:len(fileNameSep)-1]
	// 文件名
	fileInfo.Name = php2go.Implode(".", fileNameSep)
	fileInfo.TokenName = strconv.FormatInt(php2go.Microtime(), 10)
	// 文件大小
	fileInfo.Size = strconv.FormatInt(header.Size, 10)
	fileSha1, err := php2go.Sha1FileSrc(file)
	fileInfo.Hash = fileSha1
	sha1Arr := php2go.Split(fileSha1, 4)
	// 文件路径
	fileInfo.Path = "./uploads/" + php2go.Implode("/", sha1Arr) + "/"
	err = os.MkdirAll(fileInfo.Path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fileInfo.Uri = fileInfo.Path + fileInfo.TokenName + "." + fileInfo.Suffix
	php2go.Dump(fileInfo)
	out, err := os.OpenFile(fileInfo.Uri, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fileInfo, err
	}
	defer out.Close()
	io.Copy(out, file)
	return fileInfo, nil
}
