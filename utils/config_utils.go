package utils

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"hexo_editor/entity"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// ConfInit 配置文件初始化
var ConfInit string = "server:\n  # 后端运行端口\n  port: 7070\n  # hexo的根目录（此处用于限制可访问路径，比如说填写：'/data/hexo'，那么就只能访问/编辑该目录下的文件\n  hexo_root: /data/hexo\n  # 登陆用户名\n  username: xiamo\n  # 登陆密码\n  password: xiamo"

func LoadConfig(conf *entity.Config) []byte {
	var rootPath, _ = filepath.Abs(path.Dir(os.Args[0]))
	path := rootPath + "/conf.yml"
	file, err := os.Open(path)
	if err != nil {
		GenConf(rootPath)
	}
	defer file.Close()
	fileContent, _ := ioutil.ReadAll(file)
	errC := yaml.Unmarshal(fileContent, &conf)
	if errC != nil {
		log.Error("加载配置文件出错！")
		os.Exit(0)
	}
	return fileContent
}

func GenConf(path string) {
	defer os.Exit(0)
	log.Error("请在根目录中找到生成的配置文件，并将必要信息填写完整后重新运行！")

	// 配置文件
	config, _ := os.Create(path + "/conf.yml")
	config.Close()
	configFile, _ := os.OpenFile(path + "/conf.yml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	configFile.Write([]byte(ConfInit))
}
