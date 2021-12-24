package controllers

import (
	"github.com/kataras/iris/v12"
	"hexo_editor/global"
	service "hexo_editor/services"
	"hexo_editor/utils"
)

func Login(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	var user utils.User
	ctx.ReadQuery(&user)
	global.Result(ctx, 200, "success", service.Login(user))
}