package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type  TempInfo struct{
	ChunkSize int64 `json:"ChunkSize"`
	FileHash string `json:"FileHash"`
	FileId string `json:"FileId"`
	FileName string `json:"FileName"`
	FileSize int64 `json:"FileSize"`
	Label string `json:"Label"`
}

//根据传入路径判断文件或目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	//file, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
	}
	//beego.Debug(file.IsDir())
	return true
}

//根据传入路径判断文件是否是目录
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) && file.IsDir() {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
	}
	return file.IsDir()
}

//根据传入路径判断是否是文件
func IsFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) && !file.IsDir() {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
	}
	return !file.IsDir()
}

//根据传入路径判断文件是否可写
func IsWritable(path string) bool {
	return true
}

//获取路径对应的文件
func GetFile(path string) os.FileInfo {
	file, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) && !file.IsDir() {
			return file
		}
		if os.IsNotExist(err) {
			return nil
		}
	}
	if !file.IsDir() {
		return file
	} else {
		return nil
	}
}

//获取文件大小
func GetFileSize(infoPath string) int64 {
	if fileInfo, err := os.Stat(infoPath); err == nil {
		return fileInfo.Size()
	} else {
		return 0
	}
}

//删除文件
func DeleteFile(fileName string) error {
	if !IsFile(fileName) {
		return errors.New("file isn't exist")
	}
	if err := os.Remove(fileName); err != nil {
		fmt.Println("DeleteFile: " + err.Error())
		return err
	}
	return nil
}

//获取文件中的json数据
func GetJsonFileInfo(infoPath string) (temp TempInfo, n int64)  {
	tempInfoFile, err := os.Open(infoPath)
	defer tempInfoFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		//this.ReturnFailedJson(err, "Failed to find file!")
	}
	tempInfoSize := GetFileSize(infoPath)
	if tempInfoSize == 0 {
		return TempInfo{}, 0
	}
	data := TempInfo{}
	var info = make([]byte, tempInfoSize)
	if _, err := tempInfoFile.Read(info); err == nil {
		if err := json.Unmarshal(info, &data); err != nil {
			fmt.Println(err.Error())
			return TempInfo{}, 0
		}
	} else {
		fmt.Println(err.Error())
		return TempInfo{}, 0
	}
	return data, tempInfoSize
}