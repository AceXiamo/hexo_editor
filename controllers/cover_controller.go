package controllers

import (
	"github.com/kataras/iris/v12"
	"hexo_editor/global"
	service "hexo_editor/services"
)

// RandomImg
// @Description: 随机获取一张封面图
// @param ctx
func RandomImg(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	// 写入图片 字节数组
	ctx.Write(service.RandomImg())
}