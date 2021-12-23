package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// ConfInit 配置文件初始化
var ConfInit string = "{\n    \"port\": 7070,\n    \"hexo_root\": \"/data/hexo\",\n    \"username\": \"用户名\",\n    \"password\": \"密码\"\n}"

func LoadConfig(conf interface{}) {
	var rootPath, _ = filepath.Abs(path.Dir(os.Args[0]))
	path := rootPath + "/conf.json"
	file, err := os.Open(path)
	if err != nil {
		GenConf(path)
	}
	defer file.Close()
	fileContent, _ := ioutil.ReadAll(file)
	errC := json.Unmarshal(fileContent, &conf)
	if errC != nil {
		log.Error("加载配置文件出错！")
		os.Exit(0)
	}
}

func GenConf(path string) {
	defer os.Exit(0)
	log.Error("请在根目录中找到生成的配置文件，并将必要信息填写完整后重新运行！")
	file, _ := os.Create(path)
	file.Close()
	fil, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	fil.Write([]byte(ConfInit))
}
