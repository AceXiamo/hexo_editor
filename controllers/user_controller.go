package controllers

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"hexo_editor/global"
	service "hexo_editor/services"
	"hexo_editor/utils"
)

func Login(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	var user utils.User
	ctx.ReadQuery(&user)
	log.Info("[" + ctx.GetHeader("X-Real-Ip") + "] login...")
	global.Result(ctx, 200, "success", service.Login(user))
}