package route

import (
	"FILClient/web/controllers"
	"FILClient/web/controllers/env"
	"FILClient/web/controllers/user"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {

	mvc.New(app.Party("/")).Handle(controllers.NewTestController())
	main := app.Party("/", env.CrsAuth()).AllowMethods(iris.MethodOptions)
	{
		v1 := main.Party("/v1")
		{
			v1.Post("/admin/login", user.UserLogin)
			v1.PartyFunc("/admin", func(admin iris.Party) {
				admin.Use(user.JwtHandler().Serve, user.AuthToken) //登录验证
				admin.Get("/logout", user.UserLogout).Name = "退出"
			})
		}
	}

}
