package main

import (
	"github.com/kataras/iris/v12"
	"hexo_editor/global"
	"hexo_editor/router"
	"strconv"
)

func main() {
	// 初始化全局配置
	global.InitGlobalConfig()
	app := iris.New()
	app.ConfigureHost()
	// init router
	router.InitRouter(app)
	// listen port
	port := ":" + strconv.Itoa(global.Conf.Server.Port)
	app.Listen(port)
}
