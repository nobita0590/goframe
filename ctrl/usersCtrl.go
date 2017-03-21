package ctrl

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/nobita0590/goframe/models"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"errors"
	"github.com/kataras/iris"
)

type (
	UsersCtrl struct {
		mainCtrl
	}
)

func (this UsersCtrl) GetOne(ctx *iris.Context)  {
	user := models.Users{}
	if err := ctx.ReadForm(&user);err == nil{
		if err := user.Get();err == nil {
			ctx.JSON(iris.StatusOK,user)
		}else{
			this.ErrorNotFound(ctx,err)
		}
	}else{
		this.ErrorBadRequest(ctx,err)
	}
}

func (this UsersCtrl) Get(ctx *iris.Context)  {
	f := models.UsersFilter{}
	if err := ctx.ReadForm(&f);err == nil{
		user := models.Users{}
		if users,err := user.GetList(f);err == nil {
			ctx.JSON(iris.StatusOK,users)
		}else{
			this.ErrorBadRequest(ctx,err)
		}
	}else{
		this.ErrorBadRequest(ctx,err)
	}
}

func (this UsersCtrl) Options(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK,bson.M{"status":"Ok"})
}

func (this UsersCtrl) OptionsBy(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK,bson.M{"status":"Ok"})
}

func (this UsersCtrl) Post(ctx *iris.Context){
	user := models.Users{}
	if err := ctx.ReadJSON(&user);err == nil{
		fmt.Println(user)
		if err := user.Insert();err == nil{
			ctx.JSON(iris.StatusOK,user)
		}else{
			this.ErrorInternalServer(ctx,err)
		}
	}else{
		this.ErrorBadRequest(ctx,err)
	}
}

func (this UsersCtrl) Put(ctx *iris.Context){
	user := models.Users{}
	if err := ctx.ReadJSON(&user);err == nil{
		if err := user.Update();err == nil{
			ctx.JSON(iris.StatusOK,user)
		}else{
			this.ErrorInternalServer(ctx,err)
		}
	}else{
		this.ErrorBadRequest(ctx,err)
	}
}

func (this UsersCtrl) Delete(ctx *iris.Context){
	id := ctx.Param("id")
	if bson.IsObjectIdHex(id) {
		user := models.Users{Id:bson.ObjectIdHex(id)}
		if err := user.Delete();err == nil {
			ctx.JSON(iris.StatusOK,iris.Map{"Id":user.Id})
		}else{
			this.ErrorInternalServer(ctx,err)
		}
	}else{
		this.ErrorBadRequest(ctx,errors.New("params is not valid"))
	}
}