package controllers

import (
	"github.com/kataras/iris/v12"
	"hexo_editor/global"
	service "hexo_editor/services"
)

// RandomImg
// @Description: 随机封面图 - 返回url
// @param ctx
func RandomImg(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	ctx.JSON(service.RandomImg())
}

// RandomImgByte
// @Description: 随机封面图 - 返回byte
// @param ctx
func RandomImgByte(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	// 写入图片 字节数组
	ctx.Write(service.RandomImgByte())
}

// AllImg
// @Description: 获取所有封面图
// @param ctx
func AllImg(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	// 写入图片 字节数组
	ctx.JSON(service.AllImg())
}