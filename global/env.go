package global

import (
	"github.com/kataras/iris/v12"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"hexo_editor/entity"
	"hexo_editor/utils"
)

// Conf 用于保存全局配置
var Conf entity.Config

func ReloadConf()  {
	utils.LoadConfig(&Conf)
}

func InitGlobalConfig() {
	defer log.Info("Config Init Success！")
	// log样式配置
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())

	// 加载配置文件
	utils.LoadConfig(&Conf)
	// 初始化用户
	utils.InitUser(Conf)
	// 加载数据库
	utils.DbInit()
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
