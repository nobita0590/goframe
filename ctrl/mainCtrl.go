package ctrl

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
)

type (
	mainCtrl struct {
		
	}
)

func (this mainCtrl) General(ctx *iris.Context)  {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	ctx.SetHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.SetHeader("Access-Control-Max-Age","1000")
	ctx.SetHeader("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	ctx.Next()
}

func (this mainCtrl) ErrorBadRequest(ctx *iris.Context,err error)  {
	ctx.JSON(iris.StatusBadRequest,bson.M{"error":err.Error()})
}

func (this mainCtrl) ErrorInternalServer(ctx *iris.Context,err error)  {
	ctx.JSON(iris.StatusInternalServerError,bson.M{"error":err.Error()})
}

func (this mainCtrl) ErrorNotFound(ctx *iris.Context,err error)  {
	ctx.JSON(iris.StatusNotFound,bson.M{"error":err.Error()})
}