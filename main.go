package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"github.com/nobita0590/goframe/route"
	"github.com/nobita0590/goframe/models/db"
)



func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	route.Route(app)
	db.MyDb.Init()

	app.Listen(":88")
}
