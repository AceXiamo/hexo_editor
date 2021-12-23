package main

import (
	"github.com/kataras/iris/v12"
	"hexo_edit/global"
	"hexo_edit/router"
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
	port := ":" + strconv.Itoa(global.Config.Port)
	app.Listen(port)
}
