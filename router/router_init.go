package router

import (
	"github.com/kataras/iris/v12"
	sys "hexo_edit/controllers"
)

func InitRouter(app *iris.Application) {
	app.Get("/login", sys.Login)


	animeCom := app.Party("/hexo")
	{
		animeCom.Get("/files", sys.Files)
		animeCom.Get("/readFile", sys.ReadFile)
		animeCom.Post("/saveFile", sys.SaveFile)
		animeCom.Post("/removeFile", sys.RemoveFile)
		animeCom.Post("/createFile", sys.CreateFile)
	}
}
