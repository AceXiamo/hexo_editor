package services

import (
	"hexo_editor/global"
	"hexo_editor/utils"
)

// Login
// @Description: 登陆
// @param user
// @return string
func Login(user utils.User) map[string]interface{} {
	u := utils.FindUser(user.UserName)
	if (u == utils.User{} || u.PassWord != user.PassWord) {
		panic("账号或密码错误！")
	}
	code := utils.GenMd5(user.UserName, user.PassWord)
	var res map[string]interface{}
	res = make(map[string]interface{})
	res["code"] = code
	res["ali"] = global.Conf.AliConfig
	return res
}
