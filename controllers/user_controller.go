package controllers

import (
	"github.com/kataras/iris/v12"
	"hexo_edit/global"
	service "hexo_edit/services"
	"hexo_edit/utils"
)

func Login(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	var user utils.User
	ctx.ReadQuery(&user)
	global.Result(ctx, 200, "success", service.Login(user))
}