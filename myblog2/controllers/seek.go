package controllers

import (
	"github.com/astaxie/beego/orm"
	"math"
	"myblog2/models"
)

type SeekControllers struct {
	BaseControllers
}
// @router /seek [get]
func(this*SeekControllers) Seek() {
	keyboard:=this.GetStrNotNil("keyboard","搜索内容不能为空")
	PageIndex,err:=this.GetInt("PageIndex")
	if err!=nil||PageIndex<1{
		PageIndex=1
	}
	var art []models.Article
	PageSize:=10
	PageCount,_:=orm.NewOrm().QueryTable("Article").RelatedSel("User","TwoArticleType").Filter("ArtiName__icontains",keyboard).OrderBy("-Atime").Count()
	PageEnd:= int(math.Ceil(float64(PageCount)/float64(PageSize)))
	orm.NewOrm().QueryTable("Article").RelatedSel("User","TwoArticleType").Filter("ArtiName__icontains",keyboard).OrderBy("-Atime").Limit(PageSize,(PageIndex-1)*PageSize).All(&art)
	this.ShowPages(PageIndex,PageEnd)


	this.Data["keyboard"]=keyboard
	this.Data["pageEnd"]=PageEnd
	this.Data["pageIndex"]=PageIndex
	this.Data["articles"]=art
	this.Layout="layout.html"
	this.TplName="list4.html"




}