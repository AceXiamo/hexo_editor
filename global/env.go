package global

import (
	"github.com/kataras/iris/v12"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"hexo_edit/utils"
)

// 用户
var user map[string]string

type ServerConfig struct {
	Port     int `json:"port,omitempty"`
	HexoRoot string `json:"hexo_root,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// Config 用于保存全局配置
var Config ServerConfig

func InitGlobalConfig() {
	defer log.Info("配置加载完毕！")
	// log样式配置
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
	// 加载配置文件
	utils.LoadConfig(&Config)
	// 初始化用户
	utils.InitUser()
}

func Result(ctx iris.Context, code int, msg interface{}, data interface{}) {
	res := iris.Map{
		"code": code,
		"msg":  msg,
	}
	if msg == "请先登陆！" {
		res["code"] = 403
	}
	if code == 500 {
		log.Error("failed by ["+ctx.GetHeader("X-Real-Ip")+"]: ", msg)
	} else if code == 200 {
		res["data"] = data
	}
	ctx.JSON(res)
}

func ErrorHandle(ctx iris.Context) {
	if r := recover(); r != nil {
		Result(ctx, 500, r, nil)
	}
}
