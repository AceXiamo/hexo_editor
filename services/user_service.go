package services

import (
	"hexo_edit/utils"
)

// Login
// @Description: 登陆
// @param user
// @return string
func Login(user utils.User) string {
	code := ""
	u := utils.FindUser(user.UserName)
	if (u == utils.User{} || u.PassWord != user.PassWord) {
		panic("账号或密码错误！")
	}
	code = utils.GenMd5(user.UserName, user.PassWord)
	return code
}
