package router

import (
	"github.com/kataras/iris/v12"
	sys "hexo_editor/controllers"
)

func InitRouter(app *iris.Application) {
	app.Get("/login", sys.Login)

	hexo := app.Party("/hexo")
	{
		// 文件列表
		hexo.Get("/files", sys.Files)
		// 读取博客内容
		hexo.Get("/readFile", sys.ReadFile)
		// 保存博客
		hexo.Post("/saveFile", sys.SaveFile)
		// 删除博客
		hexo.Post("/removeFile", sys.RemoveFile)
		// 创建博客
		hexo.Post("/createFile", sys.CreateFile)
	}

	cover := app.Party("/cover")
	{
		// 随机封面图 - 返回url
		cover.Get("/get", sys.RandomImg)
		// 随机封面图 - 返回byte
		cover.Get("/byte", sys.RandomImgByte)
		// 获取所有封面图
		cover.Get("/all", sys.AllImg)
	}
}
