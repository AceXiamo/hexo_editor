package services

import (
	"gopkg.in/yaml.v3"
	"hexo_editor/utils"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
)

// 从配置文件中读取的图片
var images []string

func RandomImg() string {
	url := ""
	loadImg()
	// 随机从 images 中取出一张图片
	index  := rand.Intn(len(images))
	url = images[index]
	return url
}

func RandomImgByte() []byte {
	url := ""
	loadImg()
	// 随机从 images 中取出一张图片
	index  := rand.Intn(len(images))
	url = images[index]
	return utils.UrlToByte(url)
}

func AllImg() []string {
	loadImg()
	return images
}

func loadImg()  {
	// 判断 images 中是否已经加载到了图片，没图就重新读取
	if len(images) < 1 {
		// 先读配置
		var rootPath, _ = filepath.Abs(path.Dir(os.Args[0]))
		path := rootPath + "/cover.yml"
		file, err := os.Open(path)
		if err != nil {
			utils.GenCoverYml()
			panic("Error!!!")
		}
		defer file.Close()
		fileContent, _ := ioutil.ReadAll(file)
		errC := yaml.Unmarshal(fileContent, &images)
		if errC != nil {
			panic("Error!!!")
		}
	}
}

func editImg(index int, url string){

}