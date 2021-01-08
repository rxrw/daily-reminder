package main

import (
	"iuv520/daily-reminder/controllers"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"github.com/kataras/iris/v12/mvc"
)

func main() {

	app := iris.Default()

	tmpl := iris.HTML("./views", ".html")

	tmpl.Reload(true)

	tmpl.Delims("[{[", "]}]")

	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	mvc.Configure(app.Party("/script"), mvcScript)

	mvc.Configure(app.Party("/user"), mvcUser)

	app.Listen("0.0.0.0:9000")
}

func mvcScript(app *mvc.Application) {
	auth := basicauth.Default(map[string]string{
		os.Getenv("BASIC_USER"): os.Getenv("BASIC_PASS"),
	})
	app.Router.Use(auth)
	app.Handle(new(controllers.Script))
}

func mvcUser(app *mvc.Application) {
	app.Handle(new(controllers.User))
}
