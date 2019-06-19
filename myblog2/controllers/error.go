package controllers

import (
	"github.com/astaxie/beego/logs"
	"myblog2/syserror"
)

type ErrorControllers struct {
	BaseControllers
}
// ajax:{code:,msg:,url:}
func(this*ErrorControllers)Error404(){
	if this.IsAjax(){
		this.jsonErr(syserror.Error404{})
	}
	this.Layout="layout.html"
	this.TplName="error/404.html"
}
// ajax:{code:,msg:,resion:error}
func(this*ErrorControllers)Error500(){


	err,ok:=this.Data["error"].(error)
	if !ok{
		err=syserror.NewErr("未知错误",nil)

	}
	serr,ok:=err.(syserror.Error)
	if !ok{
		serr=syserror.NewErr(err.Error(),nil)
		}
	if serr.ReasonError()!=nil{
		logs.Info(serr.Error(),serr.ReasonError())
	}
	if this.IsAjax(){
	this.jsonErr(serr)
	}else {
		this.Data["errmsg"]=serr.Error()
	}
	this.Layout="layout.html"
	this.TplName="error/500.html"
}
func (this*ErrorControllers)jsonErr(serr syserror.Error)  {
	this.Ctx.Output.Status=200
	this.Data["json"]=map[string]interface{}{
		"code":serr.Code(),
		"msg":serr.Error(),
	}
	this.ServeJSON()
}