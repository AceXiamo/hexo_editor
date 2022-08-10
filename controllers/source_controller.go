package controllers

import (
	"github.com/kataras/iris/v12"
	"hexo_editor/entity"
	"hexo_editor/global"
	service "hexo_editor/services"
	"strconv"
)

// SourceByPage
//  @Description:
//  @param ctx
func SourceByPage(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	//utils.Auth(ctx.GetHeader("Auth"))
	page, _ := strconv.Atoi(ctx.URLParam("page"))
	size, _ := strconv.Atoi(ctx.URLParam("size"))

	global.Result(ctx, 200, "success", service.SourceByPage(page, size))
}

// SaveSource
//  @Description:
//  @param ctx
func SaveSource(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	//utils.Auth(ctx.GetHeader("Auth"))
	var json entity.Sources
	err := ctx.ReadJSON(&json)
	if err != nil {
		return
	}
	global.Result(ctx, 200, "success", service.SaveSource(&json))
}

// DelSource
//  @Description:
//  @param ctx
func DelSource(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	//utils.Auth(ctx.GetHeader("Auth"))
	global.Result(ctx, 200, "success", service.DelSource(ctx.URLParam("id")))
}

func UpdateSource(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	//utils.Auth(ctx.GetHeader("Auth"))
	var json entity.Sources
	err := ctx.ReadJSON(&json)
	if err != nil {
		return
	}
	global.Result(ctx, 200, "success", service.UpdateSource(&json))
}
