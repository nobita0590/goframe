package route

import (
	"gopkg.in/kataras/iris.v6"
	"eternal/iris/ctrl"
)

func Route(app *iris.Framework) {
	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusNotFound, "<h1>Custom not found handler </h1>")
	})

	app.Get("/myfiles/*file", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "Hello, the dynamic path after /myfiles is:<br/> <b>"+ctx.Param("file")+"</b>")
	})
	oAuth2(app)
	users(app)
}

func oAuth2(app *iris.Framework)  {
	oAuthCtrl := ctrl.OauthCtrl{}
	oAuthRoute := app.Party("/oauth")

	oAuthRoute.Get("/authorize",oAuthCtrl.Authorize)
	oAuthRoute.Get("/token",oAuthCtrl.Token)
}

func users(app *iris.Framework)  {
	userCtrl := ctrl.UsersCtrl{}
	userRoute := app.Party("/users", userCtrl.General)
	{
		userRoute.Options("/",userCtrl.Options)
		userRoute.Options("/:id",userCtrl.OptionsBy)
		userRoute.Delete("/:id",userCtrl.Delete)
		userRoute.Get("/",userCtrl.GetOne)
		userRoute.Get("/list",userCtrl.Get)
		userRoute.Post("/",userCtrl.Post)
		userRoute.Put("/",userCtrl.Put)
	}

}
