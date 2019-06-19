package controllers

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"myblog2/models"
)

type ContentControllers struct {
	BaseControllers
}
// @router /content [get]
func (this *ContentControllers) ShowContent(){
	id,err:=this.GetInt("id")
	if err!=nil{
		this.Abort500(errors.New("参数错误"))
	}
	artic:=this.GetContent(id)

	count,_:=orm.NewOrm().QueryTable("Praise").Filter("ArticleId",id).Count()
	this.ShowToplist(0,artic.TwoArticleType.Id)
	this.Data["count"]=count
	this.Data["article"]=artic
	this.Layout="layout.html"
	this.TplName="info.html"
}
// @router /Praise [post]
func (this *ContentControllers)Praise(){
	artId,_:=this.GetInt("articleId")
	user:=this.RtUser()
	if user==nil  {
		this.Abort500(errors.New("请登入"))
	}

	var Praise models.Praise
	o:=orm.NewOrm()
	err:=o.QueryTable("Praise").Filter("UserId",user.(models.User).Id).Filter("UserId",user.(models.User).Id).Filter("ArticleId",artId).One(&Praise)
	if err==nil{
		this.Abort500(errors.New("已经点过赞了"))
	}
	Praise.UserId=user.(models.User).Id
	Praise.ArticleId=artId

	o.Insert(&Praise)
	this.JsonOk("点赞完成","")
}