package controllers

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"myblog2/models"
)

type IndexController struct {
	BaseControllers
}

// @router / [get]
func (this *IndexController) ShowIndex() {
	var articles []models.Article
	orm.NewOrm().QueryTable("Article").RelatedSel("TwoArticleType","User").Limit(10).OrderBy("-Atime").All(&articles)



	this.ShowToplist(-1,-1)
	this.Data["articles"]=articles
	this.Layout="layout.html"
	this.TplName="index.html"
}
// @router /about.html [get]
func (this *IndexController) ShowAbout() {

	//this.Data["ttle"]="关于我"
	this.Layout="layout.html"
	this.TplName="about.html"
}
// @router /list.html [get]
func (this *IndexController) ShowDiary() {
	//this.Data["ttle"]="心情随笔"
	this.Layout="layout.html"
	this.TplName="list.html"
}
// @router /list [get]
func (this *IndexController) ShowArticle() {
	//this.Data["ttle"]="技术分享"

	TwoTypeId,err:=this.GetInt("towTypeId")
	if err!=nil{
		TwoTypeId=-1
	}
	TypeId,err:=this.GetInt("typeId")
	if err!=nil{
		this.Abort500(errors.New("获取类型失败"))
	}
	PageIndex,err:=this.GetInt("PageIndex")
	if err!=nil||PageIndex<1{
		PageIndex=1
	}
	articles,twoTypes,oneType:=this.GetArticles(TypeId,TwoTypeId,PageIndex)
	this.ShowToplist(TypeId,TwoTypeId)

	this.Data["oneType"]=oneType
	this.Data["twoArtType"]=twoTypes
	this.Data["articles"]=articles
	this.Layout="layout.html"
	this.TplName="list2.html"
}
