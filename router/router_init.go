package router

import (
	"github.com/kataras/iris/v12"
	sys "hexo_editor/controllers"
)

func InitRouter(app *iris.Application) {
	app.Get("/login", sys.Login)

	app.Get("/saveSKey", sys.SaveSKey)
	
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

	pixiv := app.Party("/pixiv")
	{
		// 获取插画作者信息
		pixiv.Get("/info", sys.PixivInfo)
		// 获取插画作者信息
		pixiv.Get("/toByte", sys.GetImageByte)
		// 分页查询插画
		pixiv.Get("/list", sys.ImageByPage)
		// 保存插画
		pixiv.Post("/save", sys.SaveImage)
		// 根据id查询插画信息
		pixiv.Get("/imageInfo", sys.ImageInfo)
		// 根据id删除插画
		pixiv.Post("/del", sys.DelImage)
	}
}
