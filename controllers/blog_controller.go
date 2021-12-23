package controllers

import (
	"github.com/kataras/iris/v12"
	"hexo_edit/global"
	service "hexo_edit/services"
	"hexo_edit/utils"
)

type JsonData struct {
	Url      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
	Content  string `json:"content"`
}

// Files
// @Description: 获取文件列表
// @param ctx
func Files(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	utils.Auth(ctx.GetHeader("Auth"))
	var json JsonData
	ctx.ReadQuery(&json)
	global.Result(ctx, 200, "success", service.GetFile(json.Url))
}

// ReadFile
// @Description: 读取文件
// @param ctx
func ReadFile(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	utils.Auth(ctx.GetHeader("Auth"))
	var json JsonData
	ctx.ReadQuery(&json)
	global.Result(ctx, 200,"success",  service.ReadFile(json.Url))
}

// SaveFile
// @Description: 保存文件
// @param ctx
func SaveFile(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	utils.Auth(ctx.GetHeader("Auth"))
	var json JsonData
	ctx.ReadJSON(&json)
	global.Result(ctx, 200,"success",  service.SaveFile(json.Url, json.Content))
}

// RemoveFile
// @Description: 删除文件
// @param ctx
func RemoveFile(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	utils.Auth(ctx.GetHeader("Auth"))
	var json JsonData
	ctx.ReadJSON(&json)
	global.Result(ctx, 200,"success",  service.RemoveFile(json.Url))
}

// CreateFile
// @Description: 创建文件
// @param ctx
func CreateFile(ctx iris.Context) {
	defer global.ErrorHandle(ctx)
	utils.Auth(ctx.GetHeader("Auth"))
	var json JsonData
	ctx.ReadJSON(&json)
	global.Result(ctx, 200,"success",  service.CreateFile(json.Url, json.Content))
}
