package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"math"
	"myblog2/models"
	"myblog2/syserror"
)

const SESSION_USER_KEY  = "SESSION_USER_KEY"
type BaseControllers struct {
	beego.Controller
	User models.User
	Islogin bool
}

func(this*BaseControllers)Prepare(){
	//user,ok:=this.GetSession(SESSION_USER_KEY).(models.User)
	//this.Islogin=false
	//if ok{
	//	this.User=user
	//	this.Islogin=true
		//this.Data["user"]=this.User

	//}
	var articleTypes []models.ArticleType
	var twoArticleTypes []models.TwoArticleType
	o:=orm.NewOrm()
	o.QueryTable("ArticleType").All(&articleTypes)
	o.QueryTable("TwoArticleType").RelatedSel("ArticleType").All(&twoArticleTypes)
	count,_:=o.QueryTable("Article").Count()
	this.Data["count"]=count
	this.Data["TwoArticleTypes"]=twoArticleTypes
	this.Data["ArticleTypes"]=articleTypes
	//this.Data["isLogin"]=this.Islogin
}
func(this*BaseControllers)Abort500(err error){
	this.Data["error"]=err
	this.Abort("500")
}
func(this*BaseControllers)GetStrNotNil(key,msg string)string{
	str:=this.GetString(key)
	if len(str)==0{
		this.Abort500(errors.New(msg))
	}
	return str
}
func(this*BaseControllers)MustLogin(){
	if !this.Islogin{
		this.Abort500(syserror.ErrorNoLogin{})
	}
}
type H map[string]interface{}
func (this*BaseControllers)JsonOk(msg,url string)  {

	this.Data["json"]=map[string]interface{}{
		"code":0,
		"msg":msg,
		"url":url,
	}
	this.ServeJSON()
}
func (this*BaseControllers)JsonOkH(msg string,data H)  {
	data["code"]=0
	data["msg"]=msg
	this.Data["json"]=data
	this.ServeJSON()
}
func (this*BaseControllers)IsAdmin()bool{
	//beego.Info(this.User.Adminrivilege)
	if this.User.Adminrivilege==0 {

		return false
	}

	return true
}
func (this*BaseControllers)UUID()string{
	u,err:=uuid.NewV4()
	if err!=nil	{
		this.Abort500(syserror.NewErr("UUID错误",err))
	}
	return u.String()
}
func (this*BaseControllers)UpImgOk(addr string){
	mp:=make(map[string]interface{})
	mp["errno"]=0
	mp["data"]=[]string{"http://www.byxiaobai.cn/static/fdfs/"+addr}
	this.Data["json"]=mp
	this.ServeJSON()
}
func (this*BaseControllers)GetArticles(TypeId ,TwoTypeId,PageIndex int)(articles []models.Article,TwoArticleTypes []models.TwoArticleType,OneType models.ArticleType ){
	o:=orm.NewOrm()
	//每一页显示几条数据
	PageSize:=10
	var PageCount int64
	if TwoTypeId==-1{
		PageCount,_= o.QueryTable("Article").RelatedSel("TwoArticleType","User").Filter("TwoArticleType__ArticleType__Id",TypeId).Count()

		o.QueryTable("Article").RelatedSel("TwoArticleType","User").Filter("TwoArticleType__ArticleType__Id",TypeId).OrderBy("-Atime").Limit(PageSize,(PageIndex-1)*PageSize).All(&articles)
	}else {
		PageCount,_= o.QueryTable("Article").RelatedSel("TwoArticleType","User").Filter("TwoArticleType__Id",TwoTypeId).Count()

		o.QueryTable("Article").RelatedSel("TwoArticleType","User").Filter("TwoArticleType__Id",TwoTypeId).OrderBy("-Atime").Limit(PageSize,(PageIndex-1)*PageSize).All(&articles)
	}
	//计算总页数
	PageEnd:= int(math.Ceil(float64(PageCount)/float64(PageSize)))
	if PageEnd<PageIndex{
		this.Abort500(syserror.Error404{})
	}
	o.QueryTable("TwoArticleType").RelatedSel("ArticleType").Filter("ArticleType__Id",TypeId).All(&TwoArticleTypes)
	OneType.Id=TypeId
	o.Read(&OneType)
	for  i,_:= range articles{
		o.LoadRelated(&articles[i],"ArticLabels")
	}
	this.ShowPages(PageIndex,PageEnd)
	this.Data["pageEnd"]=PageEnd
	this.Data["oneTypeId"]=OneType.Id
	this.Data["twoTypeId"]=TwoTypeId
	this.Data["pageIndex"]=PageIndex
	return
}
func (this*BaseControllers)GetContent(id int)(artic models.Article){
	o:=orm.NewOrm()
	artic.Id=id
	err:=o.QueryTable("Article").RelatedSel("User","TwoArticleType").Filter("Id",id).One(&artic)
	if err!=nil{
		this.Abort500(err)
	}
	artic.Acount+=1
	o.Update(&artic)
	o.LoadRelated(&artic,"ArticLabels")
	return
}
func (this*BaseControllers)ClickTop(TypeId ,TwoTypeId int){
	o:=orm.NewOrm()
	var artTop []models.Article
	if TypeId==-1{
		o.QueryTable("Article").RelatedSel("TwoArticleType","User").Limit(8).OrderBy("-Acount").All(&artTop)
		this.Data["artTop"]=artTop
		return
	}
	if TwoTypeId==-1{
		o.QueryTable("Article").RelatedSel("TwoArticleType","User").Filter("TwoArticleType__ArticleType__Id",TypeId).Limit(8).OrderBy("-Acount").All(&artTop)
	}else {
		o.QueryTable("Article").RelatedSel("TwoArticleType","User").Filter("TwoArticleType__Id",TwoTypeId).Limit(8).OrderBy("-Acount").All(&artTop)
	}


	this.Data["artTop"]=artTop
}
//页码显示
func(this*BaseControllers)ShowPages(PageIndex,PageEnd int){
	var pages []int
	if PageEnd<3{
		for i:=0;i<PageEnd;i++ {
			pages = append(pages, i+1)
		}
	}else if PageIndex>PageEnd-2{
		pages=[]int{PageEnd-2,PageEnd-1,PageEnd}
	}else if PageIndex<3{
			pages=[]int{1,2,3}
	}else {
		pages=[]int{PageIndex-1,PageIndex,PageIndex+1}
	}
	this.Data["pages"]=pages
}
//获取用户信息
func(this*BaseControllers)RtUser()interface{} {
	return this.User
}
//猜你喜欢
func(this*BaseControllers)Guess(){
	user,ok:=this.GetSession(SESSION_USER_KEY).(models.User)
	o:=orm.NewOrm()
	var GuessArticles []models.Article
	var Praise models.Praise
	if ok{
		o.QueryTable("Praise").Filter("UserId",user.Id).Limit(1).OrderBy("-Id").One(&Praise)
		var tmp models.Article
		o.QueryTable("Article").RelatedSel("TwoArticleType").Filter("Id",Praise.ArticleId).One(&tmp)
		o.QueryTable("Article").RelatedSel("TwoArticleType").Filter("TwoArticleType__Id",tmp.TwoArticleType.Id).Limit(4).OrderBy("-Id").All(&GuessArticles)



	}
	this.Data["Guess"]=GuessArticles
}
//推荐
func(this*BaseControllers)Recommend(TypeId ,TwoTypeId int){
	o:=orm.NewOrm()
	var RecommendArts []models.Article
	if TypeId==-1{
		o.QueryTable("Article").RelatedSel("TwoArticleType").Limit(8).All(&RecommendArts)
		this.Data["Recommend"]=RecommendArts
		return
	}
	if TwoTypeId!=-1{
		o.QueryTable("Article").RelatedSel("TwoArticleType").Filter("TwoArticleType__Id",TwoTypeId).Limit(8).All(&RecommendArts)
		this.Data["Recommend"]=RecommendArts
		return
	}else {
		o.QueryTable("Article").RelatedSel("TwoArticleType").Filter("TwoArticleType__ArticleType__Id",TypeId).Limit(8).All(&RecommendArts)
		this.Data["Recommend"]=RecommendArts
		return
	}

}
//推荐猜你喜欢点击排行
func(this*BaseControllers)ShowToplist(TypeId ,TwoTypeId int){
	this.ClickTop(TypeId ,TwoTypeId)
	this.Recommend(TypeId ,TwoTypeId)
	this.Guess()
}

