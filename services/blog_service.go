package services

import (
	log "github.com/sirupsen/logrus"
	"hexo_editor/global"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type HexoFile struct {
	FileName string    `json:"file_name,omitempty" xml:"file_name" url:"file_name"`
	FileSize *int64    `json:"file_size,omitempty" xml:"file_size" url:"file_size"`
	IsDir    *bool     `json:"is_dir,omitempty" xml:"is_dir" url:"is_dir"`
	ModTime  time.Time `json:"mod_time" xml:"mod_time" url:"mod_time"`
}

func GetFile(url string) []HexoFile {
	log.Info("列表：" + url)
	if url == "" || strings.Index(url, global.Config.HexoRoot) < 0 {
		panic("路径错误！")
	}
	var array []HexoFile
	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(url)
	if err != nil {
		panic(err)
	}
	// 获取文件，并输出它们的名字
	for _, file := range files {
		size := file.Size()
		isDir := file.IsDir()
		hFiles := &HexoFile{
			FileName: file.Name(),
			FileSize: &size,
			IsDir:    &isDir,
			ModTime:  file.ModTime(),
		}
		array = append(array, *hFiles)
	}
	return array
}

func ReadFile(url string) string {
	var msg string
	log.Info("读取：" + url)
	// 验证路径
	if url == "" || strings.Index(url, global.Config.HexoRoot) < 0 {
		panic("路径错误！")
	}
	// 只能读取.md
	if strings.Index(url, ".md") < 0 {
		panic("干坏事达咩！")
	}
	// 读取文件
	file, err := os.Open(url)
	if err != nil {
		msg = "文件路径有误..."
		log.Error(msg + ": " + url)
		panic("文件路径有误！")
	}
	defer file.Close()
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		panic("操作失败！")
	}
	return string(fileContent)
}

func SaveFile(url, content string) string {
	msg := "保存成功！"
	log.Info("保存：" + url)
	// 验证路径
	if url == "" || strings.Index(url, global.Config.HexoRoot) < 0 {
		panic("路径错误！")
	}
	file, err := os.OpenFile(url, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		msg = "文件路径有误..."
		log.Error(msg + ": " + url)
		panic("文件路径有误！")
	}
	defer file.Close()
	_, err = file.Write([]byte(content))
	if err != nil {
		msg = "保存文件失败..."
		panic("保存文件失败...")
	}
	return msg
}

func RemoveFile(url string) string {
	msg := "删除成功！"
	// 验证路径
	if url == "" || strings.Index(url, global.Config.HexoRoot) < 0 {
		panic("路径错误！")
	}
	// 只能删除.md文件
	if strings.Index(url, ".md") < 0 {
		panic("删除文件出错！")
	}
	os.Remove(url)
	return msg
}

func CreateFile(url, content string) string {
	msg := ""
	// 验证路径
	if url == "" || strings.Index(url, global.Config.HexoRoot) < 0 {
		panic("路径错误！")
	}
	file, err := os.Create(url)
	if err != nil {
		msg = "..."
		log.Error(msg + ": " + url)
		panic("操作失败！")
	}
	file.Close()
	// 创建成功之后保存
	msg = SaveFile(url, content)
	return msg
}
