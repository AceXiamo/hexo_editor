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
	index := rand.Intn(len(images))
	url = images[index]
	return url
}

func RandomImgByte() []byte {
	url := ""
	loadImg()
	// 随机从 images 中取出一张图片
	index := rand.Intn(len(images))
	url = images[index]
	return utils.UrlToByte(url)
}

func AllImg() []string {
	loadImg()
	return images
}

func EditImg(index int, url string) string {
	if index >= 0 {
		if url != "" {
			images[index] = url
		} else {
			images = append(images[:index], images[index+1:]...)
		}
	} else {
		images = append(images, url)
	}
	// 刷新 cover.yml
	refreshImg()
	return "操作成功!"
}

func refreshImg() {
	var str string
	str += "# 这里用来存放封面图\n"
	for _, url := range images {
		str += "- " + url + "\n"
	}
	var rootPath, _ = filepath.Abs(path.Dir(os.Args[0]))
	file, _ := os.OpenFile(rootPath + "/cover.yml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	defer file.Close()
	file.Write([]byte(str))
}

func loadImg() {
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
