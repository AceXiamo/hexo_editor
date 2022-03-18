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

type Shuo struct {
	Key    string `json:"key,omitempty"`
	OldKey string `json:"old_key,omitempty"`
}

func SaveSKey(ctx iris.Context)  {
	defer global.ErrorHandle(ctx)
	utils.Auth(ctx.GetHeader("Auth"))
	log.Info("[" + ctx.GetHeader("X-Real-Ip") + "] saveSKey...")
	var data Shuo
	ctx.ReadQuery(&data)
	global.Result(ctx, 200, "success", service.SaveSKey(data.Key, data.OldKey))
}