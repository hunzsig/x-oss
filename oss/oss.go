package oss

import (
	"../database"
	"../models"
	"../php2go"
	"mime/multipart"
	"os"
	"strconv"
)

/**
 * 判断所给路径文件/文件夹是否存在
 */
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
 * 分析文件，并返回文件信息
 */
func AnalysisFile(file multipart.File, header *multipart.FileHeader) (models.Files, error) {
	fileInfo := models.Files{}
	var err error
	if file == nil {
		file, err = header.Open()
		if err != nil {
			return fileInfo, err
		}
	}
	defer file.Close()

	// Content-Type
	fileInfo.ContentType = header.Header.Get("Content-Type")
	fileNameSep := php2go.Explode(".", header.Filename)
	// 后缀名
	fileInfo.Suffix = fileNameSep[len(fileNameSep)-1]
	fileNameSep = fileNameSep[:len(fileNameSep)-1]
	// 文件名
	fileInfo.Name = php2go.Implode(".", fileNameSep)
	// 文件大小
	fileInfo.Size = strconv.FormatInt(header.Size, 10)

	var bytes []byte
	bytes, err = php2go.FileByte(file)
	fileSha1, err := php2go.Sha1Bytes(bytes)
	if err != nil {
		return fileInfo, err
	}
	// 判断sha1是否已存在，已存在则获取数据返回
	fileInfoOld := models.Files{}
	database.Mysql().Connect.Where("hash = ?", fileSha1).First(&fileInfoOld)
	if fileInfoOld.Hash == fileSha1 {
		return fileInfoOld, nil
	}

	// md5名称
	fileInfo.Md5Name = php2go.Md5(fileSha1)
	// key
	fileInfo.Key = fileSha1
	// hash
	fileInfo.Hash = fileSha1
	// 文件路径
	sha1Arr := php2go.Split(fileSha1, 4)
	fileInfo.Path = "./uploads/" + php2go.Implode("/", sha1Arr) + "/"
	err = os.MkdirAll(fileInfo.Path, os.ModePerm)
	if err != nil {
		return fileInfo, err
	}
	// uri
	fileInfo.Uri = fileInfo.Path + fileInfo.Md5Name + "." + fileInfo.Suffix
	out, err := os.OpenFile(fileInfo.Uri, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fileInfo, err
	}
	defer out.Close()
	out.Write(bytes)
	// io.Copy(out, file)
	return fileInfo, nil
}
