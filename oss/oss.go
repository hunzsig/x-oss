package oss

import (
	"github.com/kataras/iris"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
	"x-oss/database"
	"x-oss/models"
	"x-oss/php2go"
)

const UploadsRoot = "./my_data/uploads/"

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
func AnalysisFile(ctx iris.Context, file multipart.File, header *multipart.FileHeader) (models.Files, error) {
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
	fileInfo.ContentType = strings.ToLower(header.Header.Get("Content-Type"))
	fileNameSep := php2go.Explode(".", header.Filename)
	// 后缀名
	fileInfo.Suffix = strings.ToLower(fileNameSep[len(fileNameSep)-1])
	fileNameSep = fileNameSep[:len(fileNameSep)-1]
	// 文件名
	fileInfo.Name = php2go.Implode(".", fileNameSep)
	// 文件大小
	fileInfo.Size = strconv.FormatInt(header.Size, 10)

	var bytes []byte
	bytes, err = php2go.FileByte(file)
	fileMd5, err := php2go.Md5Bytes(bytes)
	if err != nil {
		return fileInfo, err
	}
	fileSha1, err := php2go.Sha1Bytes(bytes)
	if err != nil {
		return fileInfo, err
	}
	fileHash := fileSha1 + fileMd5

	// md5名称
	fileInfo.Md5Name = php2go.Md5(fileMd5)
	// key
	fileInfo.Key = fileHash
	// hash
	fileInfo.Hash = fileHash
	// 文件路径
	now := time.Now().Format("2006-01-02 15:04:05")
	ymd := time.Now().Format("2006-01-02")
	min := time.Now().Format("15")
	fileInfo.Path = UploadsRoot + ymd + "/" + min + "/"
	fileInfo.Uri = fileInfo.Path + fileInfo.Md5Name + "." + fileInfo.Suffix

	// 判断表中 hash 是否已存在，已存在则获取数据返回
	fileInfoOld := models.Files{}
	database.Mysql().Connect.Where("hash = ?", fileHash).First(&fileInfoOld)
	if fileInfoOld.Hash == fileHash {
		return fileInfoOld, nil
	}
	// record db
	tx := database.Mysql().Connect.Begin()
	fileInfo.UserToken = ctx.Params().Get("user_token")
	fileInfo.FromUrl = ""
	fileInfo.CallQty = "0"
	fileInfo.CallLastTime = now
	fileInfo.CreateTime = now
	fileInfo.UpdateTime = now
	tx.Save(&fileInfo)
	defer tx.Close()

	// 保存文件
	err = os.MkdirAll(fileInfo.Path, os.ModePerm)
	if err != nil {
		tx.Rollback()
		return fileInfo, err
	}
	out, err := os.OpenFile(fileInfo.Uri, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		tx.Rollback()
		return fileInfo, err
	}
	defer out.Close()
	_, err = out.Write(bytes)
	// _, err = io.Copy(out, file)
	if err != nil {
		tx.Rollback()
		go os.Remove(fileInfo.Uri)
		return fileInfo, err
	}
	if tx.Commit().Error != nil {
		go os.Remove(fileInfo.Uri)
		return fileInfo, tx.Commit().Error
	}
	return fileInfo, nil
}
