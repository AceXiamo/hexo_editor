package services

import (
	"hexo_editor/global"
	"hexo_editor/utils"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
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
	res["shuo"] = global.Conf.Server.Shuo
	return res
}

func SaveSKey(key, oldKey string) string {
	var rootPath, _ = filepath.Abs(path.Dir(os.Args[0]))
	path := rootPath + "/conf.yml"
	file, err := os.Open(path)
	if err != nil {
		panic("操作失败!")
	}
	defer file.Close()
	fileContent, _ := ioutil.ReadAll(file)
	ctn := string(fileContent)
	ctn = strings.Replace(ctn, oldKey, key, -1)
	configFile, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	configFile.Write([]byte(ctn))
	// 保存完了重新加载配置文件
	global.ReloadConf()
	return "修改成功!"
}
