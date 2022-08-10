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
		// 编辑封面图
		cover.Post("/edit", sys.EditImg)
	}

	pixiv := app.Party("/pixiv")
	{
		// 获取插画作者信息
		pixiv.Get("/info", sys.PixivInfo)
		// 根据P站插画路径取图
		pixiv.Get("/toByte", sys.GetImageByte)
		// 分页查询插画
		pixiv.Get("/list", sys.ImageByPage)
		// 保存插画
		pixiv.Post("/save", sys.SaveImage)
		// 根据id查询插画信息
		pixiv.Get("/imageInfo", sys.ImageInfo)
		// 根据id删除插画
		pixiv.Post("/del", sys.DelImage)
		// 查询所有已保存作者信息，一次性返回不进行分页
		pixiv.Get("/auth", sys.AuthList)
	}

	source := app.Party("/source")
	{
		// 分页查询
		source.Get("/list", sys.SourceByPage)
		// 保存
		source.Post("/save", sys.SaveSource)
		// 根据id删除
		source.Post("/del", sys.DelSource)
		// 修改
		source.Post("/update", sys.UpdateSource)
	}
}
