package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"myblog2/controllers"
	"myblog2/models"
)

func init() {
	beego.InsertFilter("/admin/*",beego.BeforeExec,FnUser)
	beego.ErrorController(&controllers.ErrorControllers{})
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserControllers{})
	beego.Include(&controllers.ArticleControllers{})
	beego.Include(&controllers.ContentControllers{})
	beego.Include(&controllers.SeekControllers{})
}
var FnUser = func(c*context.Context){
	var user models.User
	User:=c.Input.Session("SESSION_USER_KEY")
	if User==nil{
		c.Redirect(302,"/")
		return
	}
	user=User.(models.User)
	if user.Adminrivilege==0{
		c.Redirect(302,"/")
	}
}