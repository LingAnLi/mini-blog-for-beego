package main

import (
	"encoding/gob"
	"myblog2/models"
	_ "myblog2/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_"myblog2/models"
)

func main() {
	beego.AddFuncMap("PrePage",PrePage)
	beego.AddFuncMap("NextPage",NextPage)
	beego.AddFuncMap("ShowI",ShowI)
	gob.Register(models.User{})
	initSession()
	beego.Run()
}
// session
func initSession()  {
	beego.BConfig.WebConfig.Session.SessionOn=true
	beego.BConfig.WebConfig.Session.SessionName="blog"
	beego.BConfig.WebConfig.Session.SessionProvider="file"

	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}
func ShowI(int2 int)int{
	return int2+1
}
func NextPage(pageIndex,pageEnd int) int {
	if pageIndex>=pageEnd {
		return pageEnd
	}
	return pageIndex+1
}
func PrePage(PageIndex int)int  {
	if PageIndex<=1{
		return 1
	}
	return PageIndex-1
}