package utils

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func UrlToByte(url string) []byte {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// GenFileToRoot
// @Description: 在根目录下创建一个文件
func GenFileToRoot(fileName string)  {
	rootPath, _ := filepath.Abs(path.Dir(os.Args[0]))
	config, _ := os.Create(rootPath + fileName)
	err := config.Close()
	if err != nil {
		panic(fileName + ": 创建失败!")
	}
}