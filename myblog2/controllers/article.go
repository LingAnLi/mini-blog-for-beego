package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/weilaihui/fdfs_client"
	"myblog2/models"
	"path/filepath"
	"strconv"
	"strings"
)

type ArticleControllers struct {
	BaseControllers
}
//@router /admin/index.html [get]
func(this*ArticleControllers) ShowIndex() {
	var twoArticleTypes []models.TwoArticleType
	o:=orm.NewOrm()
	o.QueryTable("TwoArticleType").RelatedSel("ArticleType").All(&twoArticleTypes)
	var articles []models.Article
	o.QueryTable("Article").RelatedSel("TwoArticleType","User").All(&articles)
	this.Data["article"]=articles
	this.Data["twoArtTypes"]=twoArticleTypes
	this.Layout="admin/adminlayout.html"
	this.TplName="admin/index.html"
}
//@router /admin/addType.html [get]
func(this*ArticleControllers) ShowAddType() {
	this.IsAdmin()
	var articleTypes []models.ArticleType
	var twoArticleTypes []models.TwoArticleType
	o:=orm.NewOrm()
	o.QueryTable("ArticleType").All(&articleTypes)
	o.QueryTable("TwoArticleType").RelatedSel("ArticleType").All(&twoArticleTypes)

	this.Data["articleTypes"]=articleTypes
	this.Data["twoArticleTypes"]=twoArticleTypes
	beego.Info(articleTypes)

	this.Layout="admin/adminlayout.html"
	this.TplName="admin/addType.html"
}
//添加分类
//@router /admin/addtype [get]
func(this*ArticleControllers) HanderAddType() {
	typeName1:=this.GetStrNotNil("typeName1","分类1不能为空")
	typeName2:=this.GetString("typeName2")
	beego.Info(typeName1,"||",typeName2)
	o:=orm.NewOrm()
	var ArticleTypes1 []models.ArticleType
	o.QueryTable("ArticleType").All(&ArticleTypes1)
	//添加一级分类
	if len(typeName2)==0{
		for _,art:=range ArticleTypes1{
			if art.TypeName==typeName1{
				this.Abort500(errors.New(typeName1+"一级分类已存在"))
			}
		}
		var articType models.ArticleType
		articType.TypeName=typeName1
		o.Insert(&articType)
	}
	//添加二级分类
	if len(typeName2)!=0{
		ok:=false
		for _,art:=range ArticleTypes1{
			if art.TypeName==typeName1{
				ok=true
			}
		}
		if ok {
			//一级分类存在添加二级分类
			var articType models.ArticleType
			var twoArticType models.TwoArticleType
			articType.TypeName=typeName1
			o.Read(&articType,"TypeName")
			twoArticType.TypeName=typeName2
			twoArticType.ArticleType=&articType
			o.Insert(&twoArticType)
		}else {
			this.Abort500(errors.New(typeName1+"一级分类不存在，无法添加二级分类"))
		}


	}
	//返回视图
	this.Redirect("/admin/addType.html",302)
}
//删除分类
//@router /admin/deleteType [get]
func(this*ArticleControllers)DeleteType(){
	id1,err:=this.GetInt("typeid")
	o:=orm.NewOrm()
	//删除一级分类
	if err==nil{
		var articleType models.ArticleType
		articleType.Id=id1
		o.Delete(&articleType)
	}

	id2,err:=this.GetInt("type2id")
	//删除二级分类
	if err==nil{
		var twoArticleType models.TwoArticleType
		twoArticleType.Id=id2
		o.Delete(&twoArticleType)
	}


	//返回数据
	this.Redirect("/admin/addType.html",302)
}
//显示文章标签
//@router /admin/addLable.html [get]
func(this*ArticleControllers)ShowLAble(){
	var articleLable []models.ArticLabel
	orm.NewOrm().QueryTable("ArticLabel").All(&articleLable)
	this.Data["lable"]=articleLable
	this.Layout="admin/adminlayout.html"
	this.TplName="admin/addLable.html"

}
//添加文章标签
//@router /admin/addLable [get]
func(this*ArticleControllers)AddLAble(){
	LableName:=this.GetStrNotNil("LableName","标签不能为空")
	var articleLable models.ArticLabel
	articleLable.Name=LableName
	o:=orm.NewOrm()
	_,err:=o.Insert(&articleLable)
	if err!=nil{
		this.Abort500(err)
	}
	this.Redirect("/admin/addLable.html",302)

}
//删除标签
//@router /admin/deleteLable [get]
func(this*ArticleControllers)DeleteLAble(){
	id,err:=this.GetInt("lableid")
	if err!=nil{
		this.Abort500(err)
	}
	var articleLable models.ArticLabel
	articleLable.Id=id
	_,err=orm.NewOrm().Delete(&articleLable)
	if err!=nil{
		this.Abort500(err)
	}
	this.Redirect("/admin/addLable.html",302)
}
//显示添加文章
//@router /admin/add.html	[get]
func(this*ArticleControllers)ShowAddArticle(){
	var twoArtType []models.TwoArticleType
	var articleLabel []models.ArticLabel
	o:=orm.NewOrm()
	o.QueryTable("ArticLabel").All(&articleLabel)
	o.QueryTable("TwoArticleType").RelatedSel("ArticleType").All(&twoArtType)

	this.Data["key"]=this.UUID()//文章唯一标识
	this.Data["articleLabel"]=articleLabel
	this.Data["twoArtType"]=twoArtType
	this.Layout="admin/adminlayout.html"
	this.TplName="admin/add.html"
}
//上传图片
//@router /admin/upload/:key [post]
func(this*ArticleControllers)UploadImg(){
	//获取文章唯一标识
	key:=this.Ctx.Input.Param(":key")
	if len(key)==0{

	}
	//保存数据
	conn,err:=fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
	if err!=nil{
		beego.Error("连接错误",err)
		this.Abort500(err)
	}
	f,h,err:=this.GetFile("ImgName")
	if err!=nil{
		beego.Error("获取图片错误",err)
		this.Abort500(err)
	}
	fbuf:=make([]byte,h.Size)
	f.Read(fbuf)
	ext:=filepath.Ext(h.Filename)
	res,err:=conn.UploadByBuffer(fbuf,ext[1:])
	if err!=nil{
		beego.Error("存图片错误",err)
		this.Abort500(err)
	}
	var Img models.ArticIMG
	Img.Addr=res.RemoteFileId
	Img.Key=key
	_,err=orm.NewOrm().Insert(&Img)
	if err!=nil{
		beego.Error("存图片错误",err)
		this.Abort500(err)
	}
	this.UpImgOk(Img.Addr)
}
//实现添加文章功能
//@router /admin/addArticle/:key [post]
func(this*ArticleControllers)HanderAddArticle(){
	//获取数据
	//beego.Info(this.Ctx.Request.RemoteAddr)
	key:=this.Ctx.Input.Param(":key")
	articleName:=this.GetStrNotNil("artiName","标题不能为空")
	content:=this.GetStrNotNil("content","内容不能为空")
	twoType:=this.GetStrNotNil("twoArticleName","分类不能为空")
	mylabelstr:=this.GetString("mylabel")
	//处理数据
	o:=orm.NewOrm()
	var twoArticleType models.TwoArticleType
	var article models.Article
	//保证key唯一
	err:=o.QueryTable("Article").Filter("Key",key).One(&article)
	if err==nil{
		this.Abort500(errors.New("文章保存失败,文章ID已存在"))
	}
	twoTypeId,_:=strconv.Atoi(twoType)
	twoArticleType.Id=twoTypeId
	o.Read(&twoArticleType)
	article.ArtiName=articleName
	article.Acontent=content
	article.TwoArticleType=&twoArticleType
	article.Key=key
	article.User=&this.BaseControllers.User
	//插入文章
	_,err=o.Insert(&article)
	if err!=nil{
		this.Abort500(errors.New("文章保存失败"))
	}
	//插入文章标签如果添加了
	mylabels:=strings.Split(mylabelstr,"-")
	m2m:=o.QueryM2M(&article,"ArticLabels")
	if len(mylabelstr)!=0{
		for _,labelId:=range mylabels{
			var articLabel models.ArticLabel
			ilabelId,_:=strconv.Atoi(labelId)
			articLabel.Id=ilabelId
			beego.Info(articLabel.Id)
			err:=o.Read(&articLabel)
			if err!=nil{
				this.Abort500(errors.New("xxxx"))
			}

			_,err=m2m.Add(&articLabel)
			if err!=nil{
				this.Abort500(err)
			}
		}
	}


	this.JsonOk("ok","/")

}
//查看详情
//@router /admin/content [get]
func(this*ArticleControllers)ShowArticle() {
	id:=this.GetStrNotNil("id","该文章不存在")
	var article  models.Article
	orm.NewOrm().QueryTable("Article").RelatedSel("TwoArticleType").Filter("Id",id).One(&article)
	this.Data["article"]=article
	this.Layout="admin/adminlayout.html"
	this.TplName="admin/content.html"
}